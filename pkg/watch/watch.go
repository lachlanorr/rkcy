// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package watch

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	"github.com/lachlanorr/rocketcycle/pkg/jsonutils"
	"github.com/lachlanorr/rocketcycle/pkg/mgmt"
	"github.com/lachlanorr/rocketcycle/pkg/rkcy"
	"github.com/lachlanorr/rocketcycle/pkg/rkcypb"
	"github.com/lachlanorr/rocketcycle/version"
)

type watchTopic struct {
	clusterName    string
	brokers        string
	topicName      string
	partitionCount int32
	logLevel       zerolog.Level
}

func (wt *watchTopic) String() string {
	return fmt.Sprintf("%s__%s", wt.clusterName, wt.topicName)
}

func (wt *watchTopic) consume(
	ctx context.Context,
	wg *sync.WaitGroup,
	strmprov rkcy.StreamProvider,
	cncHdlrs rkcy.ConcernHandlers,
	watchDecode bool,
) {
	groupName := rkcy.UncommittedGroupNameAllPartitions(wt.topicName)
	log.Info().Msgf("watching: %s", wt.topicName)

	kafkaLogCh := make(chan kafka.LogEvent)
	go rkcy.PrintKafkaLogs(ctx, kafkaLogCh)
	cons, err := strmprov.NewConsumer(wt.brokers, groupName, kafkaLogCh)
	if err != nil {
		log.Error().
			Err(err).
			Str("BootstrapServers", wt.brokers).
			Str("GroupId", groupName).
			Msg("Unable to NewConsumer")
		return
	}
	defer cons.Close()

	topicParts := make([]kafka.TopicPartition, wt.partitionCount)
	for i := int32(0); i < wt.partitionCount; i++ {
		topicParts[i] = kafka.TopicPartition{
			Topic:     &wt.topicName,
			Partition: i,
			Offset:    kafka.OffsetEnd,
		}
	}
	cons.Assign(topicParts)

	pjOpts := protojson.MarshalOptions{Multiline: true, Indent: "  ", EmitUnpopulated: true}

	wg.Add(1)
	for {
		select {
		case <-ctx.Done():
			log.Trace().
				Msg("watchTopic.consume exiting, ctx.Done()")
			wg.Done()
			return
		default:
			msg, err := cons.ReadMessage(100 * time.Millisecond)
			timedOut := err != nil && err.(kafka.Error).Code() == kafka.ErrTimedOut
			if err != nil && !timedOut {
				log.Error().
					Err(err).
					Msg("Error during ReadMessage")
			} else if !timedOut && msg != nil {
				log.WithLevel(wt.logLevel).
					Str("Directive", fmt.Sprintf("0x%08X", int(rkcy.GetDirective(msg)))).
					Str("TxnId", rkcy.GetTraceId(msg)).
					Int("Offset", int(msg.TopicPartition.Offset)).
					Msg(wt.topicName)
				txn := &rkcypb.ApecsTxn{}
				err := proto.Unmarshal(msg.Value, txn)
				if err == nil {
					if watchDecode {
						txnJsonDec, err := DecodeTxnOpaques(ctx, txn, cncHdlrs, &pjOpts)
						if err == nil {
							log.WithLevel(wt.logLevel).
								Msg(string(txnJsonDec))
						} else {
							log.Error().
								Err(err).
								Msgf("Failed to decodeOpaques: %s", err.Error())
						}
					} else {
						txnJson := pjOpts.Format(txn)
						log.WithLevel(wt.logLevel).
							Msg(string(txnJson))
					}
				}
			}
		}
	}
}

func getAllWatchTopics(rtPlatDef *rkcy.RtPlatformDef) []*watchTopic {
	var wts []*watchTopic
	for _, concern := range rtPlatDef.Concerns {
		for _, topic := range concern.Topics {
			tp, err := rkcy.ParseFullTopicName(topic.CurrentTopic)
			if err == nil {
				if tp.Topic == rkcy.ERROR || tp.Topic == rkcy.COMPLETE {
					wt := watchTopic{
						clusterName:    topic.CurrentCluster.Name,
						brokers:        topic.CurrentCluster.Brokers,
						topicName:      topic.CurrentTopic,
						partitionCount: topic.CurrentTopicPartitionCount,
					}
					if tp.Topic == rkcy.ERROR {
						wt.logLevel = zerolog.ErrorLevel
					} else {
						wt.logLevel = zerolog.DebugLevel
					}
					wts = append(wts, &wt)
				}
			}
		}
	}
	return wts
}

func DecodeTxnOpaques(ctx context.Context, txn *rkcypb.ApecsTxn, concernHandlers rkcy.ConcernHandlers, pjOpts *protojson.MarshalOptions) ([]byte, error) {
	txnJson, err := pjOpts.Marshal(txn)
	if err != nil {
		return nil, err
	}

	var txnTopLvl *jsonutils.OrderedMap
	err = jsonutils.UnmarshalOrdered(txnJson, &txnTopLvl)
	if err != nil {
		return nil, err
	}

	var decSteps = func(iSteps interface{}) {
		steps, ok := iSteps.([]interface{})
		if ok {
			for _, iStep := range steps {
				step, ok := iStep.(*jsonutils.OrderedMap)
				if ok {
					iConcern, ok := step.Get("concern")
					if !ok {
						continue
					}
					concern, ok := iConcern.(string)
					if !ok {
						continue
					}

					iSystem, ok := step.Get("system")
					if !ok {
						continue
					}
					systemStr, ok := iSystem.(string)
					if !ok {
						continue
					}
					systemInt32, ok := rkcypb.System_value[systemStr]
					if !ok {
						continue
					}
					system := rkcypb.System(systemInt32)

					iCommand, ok := step.Get("command")
					if !ok {
						continue
					}
					command, ok := iCommand.(string)
					if !ok {
						continue
					}

					iB64, ok := step.Get("payload")
					if ok {
						b64, ok := iB64.(string)
						if ok && b64 != "" {
							argJson, err := concernHandlers.DecodeArgPayload64Json(ctx, concern, system, command, b64)
							var argOmap *jsonutils.OrderedMap
							err = jsonutils.UnmarshalOrdered(argJson, &argOmap)
							if err == nil {
								step.SetAfter("payloadDec", argOmap, "payload")
							} else {
								step.SetAfter("payloadDecErr", err.Error(), "payload")
							}
						}
					}

					iB64, ok = step.Get("instance")
					if ok {
						b64, ok := iB64.(string)
						if ok && b64 != "" {
							argJson, err := concernHandlers.DecodeInstance64Json(ctx, concern, b64)
							var argOmap *jsonutils.OrderedMap
							err = jsonutils.UnmarshalOrdered(argJson, &argOmap)
							if err == nil {
								step.SetAfter("instanceDec", argOmap, "instance")
							} else {
								step.SetAfter("instanceDecErr", err.Error(), "instance")
							}
						}
					}

					iRslt, ok := step.Get("result")
					if ok {
						rslt, ok := iRslt.(*jsonutils.OrderedMap)
						if ok {
							iB64, ok := rslt.Get("payload")
							if ok {
								b64, ok := iB64.(string)
								if ok && b64 != "" {
									resJson, err := concernHandlers.DecodeResultPayload64Json(ctx, concern, system, command, b64)
									var resOmap *jsonutils.OrderedMap
									err = jsonutils.UnmarshalOrdered(resJson, &resOmap)
									if err == nil {
										rslt.SetAfter("payloadDec", resOmap, "payload")
									} else {
										rslt.SetAfter("payloadDecErr", err.Error(), "payload")
									}
								}
							}

							iB64, ok = rslt.Get("instance")
							if ok {
								b64, ok := iB64.(string)
								if ok && b64 != "" {
									instJson, err := concernHandlers.DecodeInstance64Json(ctx, concern, b64)
									var instOmap *jsonutils.OrderedMap
									err = jsonutils.UnmarshalOrdered(instJson, &instOmap)
									if err == nil {
										rslt.SetAfter("instanceDec", instOmap, "instance")
									} else {
										rslt.SetAfter("instanceDecErr", err.Error(), "instance")
									}
								}
							}
						}
					}
				}
			}
		}
	}

	iFwSteps, ok := txnTopLvl.Get("forwardSteps")
	if ok {
		decSteps(iFwSteps)
	}
	iRevSteps, ok := txnTopLvl.Get("reverseSteps")
	if ok {
		decSteps(iRevSteps)
	}

	jsonSer, err := jsonutils.MarshalOrderedIndent(txnTopLvl, "", "  ")
	if err != nil {
		return nil, err
	}

	return jsonSer, nil
}

func WatchResultTopics(
	ctx context.Context,
	wg *sync.WaitGroup,
	strmprov rkcy.StreamProvider,
	platform string,
	environment string,
	adminBrokers string,
	cncHdlrs rkcy.ConcernHandlers,
	watchDecode bool,
) {
	defer wg.Done()

	wtMap := make(map[string]bool)

	platformCh := make(chan *mgmt.PlatformMessage)
	mgmt.ConsumePlatformTopic(
		ctx,
		wg,
		strmprov,
		platform,
		environment,
		adminBrokers,
		platformCh,
		nil,
	)

	for {
		select {
		case <-ctx.Done():
			log.Trace().
				Msg("WatchResultTopics exiting, ctx.Done()")
			return
		case platMsg := <-platformCh:
			if (platMsg.Directive & rkcypb.Directive_PLATFORM) != rkcypb.Directive_PLATFORM {
				log.Error().Msgf("Invalid directive for PlatformTopic: %s", platMsg.Directive.String())
				continue
			}

			wts := getAllWatchTopics(platMsg.NewRtPlatDef)

			for _, wt := range wts {
				_, ok := wtMap[wt.String()]
				if !ok {
					wtMap[wt.String()] = true
					go wt.consume(
						ctx,
						wg,
						strmprov,
						cncHdlrs,
						watchDecode,
					)
				}
			}
		}
	}
}

func Start(
	ctx context.Context,
	wg *sync.WaitGroup,
	strmprov rkcy.StreamProvider,
	platform string,
	environment string,
	adminBrokers string,
	cncHdlrs rkcy.ConcernHandlers,
	watchDecode bool,
) {
	log.Info().
		Str("GitCommit", version.GitCommit).
		Msg("APECS WATCH started")

	wg.Add(1)
	go WatchResultTopics(
		ctx,
		wg,
		strmprov,
		platform,
		environment,
		adminBrokers,
		cncHdlrs,
		watchDecode,
	)

	wg.Add(1)
	select {
	case <-ctx.Done():
		log.Info().
			Msg("watch stopped")
		wg.Done()
		return
	}
}
