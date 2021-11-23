// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rkcy

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	otel_codes "go.opentelemetry.io/otel/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

var nameRe = regexp.MustCompile(`^[a-zA-Z][a-zA-Z\-]{1,15}$`)

func IsValidName(name string) bool {
	return nameRe.MatchString(name)
}

type Settings struct {
	PlatformFilePath string
	ConfigFilePath   string

	OtelcolEndpoint string

	AdminBrokers    string
	ConsumerBrokers string

	HttpAddr   string
	GrpcAddr   string
	PortalAddr string

	System    System
	Topic     string
	Partition int32

	AdminPingIntervalSecs uint

	StorageTarget string

	WatchDecode bool
}

type ConcernHandlers map[string]ConcernHandler

type Platform struct {
	name          string
	environment   string
	cobraCommands []*cobra.Command
	storageInits  map[string]StorageInit

	settings  Settings
	telem     *Telemetry
	configMgr *ConfigMgr

	adminPingInterval time.Duration

	concernHandlers ConcernHandlers
	rawProducer     *RawProducer

	producerTracker  *ProducerTracker
	currentRtPlatDef *rtPlatformDef
}

func NewPlatform(
	name string,
	environment string,
) (*Platform, error) {
	if !IsValidName(name) {
		return nil, fmt.Errorf("Invalid name: %s", name)
	}
	environment = os.Getenv("RKCY_ENVIRONMENT")
	if !IsValidName(environment) {
		return nil, fmt.Errorf("Invalid RKCY_ENVIRONMENT: %s", environment)
	}

	environment = os.Getenv("RKCY_ENVIRONMENT")
	if !IsValidName(environment) {
		return nil, fmt.Errorf("Invalid RKCY_ENVIRONMENT: %s", environment)
	}

	plat := &Platform{
		name:          name,
		environment:   environment,
		cobraCommands: make([]*cobra.Command, 0),
		storageInits:  make(map[string]StorageInit),

		settings: Settings{Partition: -1},

		concernHandlers: gConcernHandlerRegistry,
	}

	return plat, nil
}

func (plat *Platform) AdminPingInterval() time.Duration {
	if plat.adminPingInterval == 0 {
		plat.adminPingInterval = time.Duration(plat.settings.AdminPingIntervalSecs) * time.Second
	}
	return plat.adminPingInterval
}

func (plat *Platform) Start() {
	prepLogging(plat.name)
	// validate all command handlers exist for each concern
	if !plat.concernHandlers.validateConcernHandlers() {
		log.Fatal().
			Msg("validateConcernHandlers failed")
	}
	plat.runCobra()
}

func (plat *Platform) Name() string {
	return plat.name
}

func (plat *Platform) Environment() string {
	return plat.environment
}

func (plat *Platform) Telem() *Telemetry {
	return plat.telem
}

func (plat *Platform) Settings() Settings {
	return plat.settings
}

func (plat *Platform) AppendCobraCommand(cmd *cobra.Command) {
	plat.cobraCommands = append(plat.cobraCommands, cmd)
}

func (plat *Platform) SetStorageInit(name string, storageInit StorageInit) {
	plat.storageInits[name] = storageInit
}

func (plat *Platform) RegisterLogicHandler(concern string, handler interface{}) {
	plat.concernHandlers.RegisterLogicHandler(concern, handler)
}

func (plat *Platform) RegisterCrudHandler(storageType string, concern string, handler interface{}) {
	plat.concernHandlers.RegisterCrudHandler(storageType, concern, handler)
}

func (plat *Platform) ConfigMgr() *ConfigMgr {
	if plat.configMgr == nil {
		log.Fatal().Msg("InitConfigMgr has not been called")
	}
	return plat.configMgr
}

func (plat *Platform) InitConfigMgr(ctx context.Context, wg *sync.WaitGroup) {
	if plat.configMgr != nil {
		log.Fatal().Msg("InitConfigMgr called twice")
	}
	plat.configMgr = NewConfigMgr(
		ctx,
		plat.settings.AdminBrokers,
		plat.name,
		plat.environment,
		wg,
	)
}

type rtPlatformDef struct {
	PlatformDef          *PlatformDef
	Hash                 string
	Concerns             map[string]*rtConcern
	Clusters             map[string]*Cluster
	AdminCluster         string
	StorageTargets       map[string]*StorageTarget
	PrimaryStorageTarget string
}

type rtConcern struct {
	Concern *Concern
	Topics  map[string]*rtTopics
}

type rtTopics struct {
	Topics                     *Concern_Topics
	CurrentTopic               string
	CurrentTopicPartitionCount int32
	CurrentCluster             *Cluster
	FutureTopic                string
	FutureTopicPartitionCount  int32
	FutureCluster              *Cluster
}

func newRtConcern(rtPlatDef *rtPlatformDef, concern *Concern) (*rtConcern, error) {
	rtConc := rtConcern{
		Concern: concern,
		Topics:  make(map[string]*rtTopics),
	}
	for _, topics := range concern.Topics {
		// verify topics only appear once
		if _, ok := rtConc.Topics[topics.Name]; ok {
			return nil, fmt.Errorf("Topic '%s' appears more than once in Concern '%s' definition", topics.Name, rtConc.Concern.Name)
		}
		rtTops, err := newRtTopics(rtPlatDef, &rtConc, topics)
		if err != nil {
			return nil, err
		}
		rtConc.Topics[topics.Name] = rtTops
	}
	return &rtConc, nil
}

func isACETopic(topic string) bool {
	return topic == string(ADMIN) || topic == string(ERROR) || topic == string(COMPLETE)
}

func newRtTopics(rtPlatDef *rtPlatformDef, rtConc *rtConcern, topics *Concern_Topics) (*rtTopics, error) {
	rtTops := rtTopics{
		Topics: topics,
	}

	pref := BuildTopicNamePrefix(rtPlatDef.PlatformDef.Name, rtPlatDef.PlatformDef.Environment, rtConc.Concern.Name, rtConc.Concern.Type)
	var ok bool

	rtTops.CurrentTopic = BuildTopicName(pref, topics.Name, topics.Current.Generation)
	rtTops.CurrentTopicPartitionCount = maxi32(1, topics.Current.PartitionCount)
	rtTops.CurrentCluster, ok = rtPlatDef.Clusters[topics.Current.Cluster]
	if !ok {
		return nil, fmt.Errorf("Topic '%s.%s' has invalid Current Cluster '%s'", rtConc.Concern.Name, topics.Name, topics.Current.Cluster)
	}

	if topics.Future != nil {
		rtTops.FutureTopic = BuildTopicName(pref, topics.Name, topics.Future.Generation)
		rtTops.FutureTopicPartitionCount = maxi32(1, topics.Future.PartitionCount)
		rtTops.FutureCluster, ok = rtPlatDef.Clusters[topics.Future.Cluster]
		if !ok {
			return nil, fmt.Errorf("Topic '%s.%s' has invalid Future Cluster '%s'", rtConc.Concern.Name, topics.Name, topics.Future.Cluster)
		}
	}

	if isACETopic(rtTops.Topics.Name) {
		if rtTops.CurrentTopicPartitionCount != 1 {
			return nil, fmt.Errorf("Topic '%s.%s' has invalid partition count current=%d", rtConc.Concern.Name, topics.Name, rtTops.CurrentTopicPartitionCount)
		}
	}

	return &rtTops, nil
}

func initTopic(topic *Concern_Topic, adminCluster string) *Concern_Topic {
	if topic == nil {
		topic = &Concern_Topic{}
	}

	if topic.Generation <= 0 {
		topic.Generation = 1
	}
	if topic.Cluster == "" {
		topic.Cluster = adminCluster
	}
	if topic.PartitionCount <= 0 {
		topic.PartitionCount = 1
	} else if topic.PartitionCount > MAX_PARTITION {
		topic.PartitionCount = MAX_PARTITION
	}

	return topic
}

func initTopics(
	topics *Concern_Topics,
	adminCluster string,
	concernType Concern_Type,
	storageTargets []*StorageTarget,
) *Concern_Topics {
	if topics == nil {
		topics = &Concern_Topics{}
	}

	topics.Current = initTopic(topics.Current, adminCluster)
	if topics.Future != nil {
		topics.Future = initTopic(topics.Future, adminCluster)
	}

	if concernType == Concern_APECS {
		topics.ConsumerPrograms = nil
		switch topics.Name {
		case "process":
			prog := &Program{
				Name:   "./@platform",
				Args:   []string{"process", "--otelcol_endpoint", "@otelcol_endpoint", "--admin_brokers", "@admin_brokers", "--consumer_brokers", "@consumer_brokers", "-t", "@topic", "-p", "@partition"},
				Abbrev: "p/@concern/@partition",
			}
			topics.ConsumerPrograms = append(topics.ConsumerPrograms, prog)
		case "storage":
			for _, stgTgt := range storageTargets {
				if stgTgt.IsPrimary {
					prog := &Program{
						Name: "./@platform",
						Args: []string{"storage", "--otelcol_endpoint", "@otelcol_endpoint", "--admin_brokers", "@admin_brokers", "--consumer_brokers", "@consumer_brokers", "-t", "@topic", "-p", "@partition", "--storage_target", stgTgt.Name},
					}
					prog.Abbrev = fmt.Sprintf("s/*%s/@concern/@partition", stgTgt.Name)
					topics.ConsumerPrograms = append(topics.ConsumerPrograms, prog)
				}
			}
		case "storage-scnd":
			for _, stgTgt := range storageTargets {
				if !stgTgt.IsPrimary {
					prog := &Program{
						Name: "./@platform",
						Args: []string{"storage-scnd", "--otelcol_endpoint", "@otelcol_endpoint", "--admin_brokers", "@admin_brokers", "--consumer_brokers", "@consumer_brokers", "-t", "@topic", "-p", "@partition", "--storage_target", stgTgt.Name},
					}
					prog.Abbrev = fmt.Sprintf("s/%s/@concern/@partition", stgTgt.Name)
					topics.ConsumerPrograms = append(topics.ConsumerPrograms, prog)
				}
			}
		}
	}

	return topics
}

func newRtPlatformDef(platDef *PlatformDef, platformName string, environment string) (*rtPlatformDef, error) {
	if platDef.Name != platformName {
		return nil, fmt.Errorf("Platform Name mismatch, '%s' != '%s'", platDef.Name, platformName)
	}
	if platDef.Environment != environment {
		return nil, fmt.Errorf("Environment mismatch, '%s' != '%s'", platDef.Environment, environment)
	}

	rtPlatDef := rtPlatformDef{
		PlatformDef:    platDef,
		Concerns:       make(map[string]*rtConcern),
		Clusters:       make(map[string]*Cluster),
		StorageTargets: make(map[string]*StorageTarget),
	}

	platJson := protojson.Format(proto.Message(rtPlatDef.PlatformDef))
	sha256Bytes := sha256.Sum256([]byte(platJson))
	rtPlatDef.Hash = hex.EncodeToString(sha256Bytes[:])

	if !rtPlatDef.PlatformDef.UpdateTime.IsValid() {
		return nil, fmt.Errorf("Invalid UpdateTime: %s", rtPlatDef.PlatformDef.UpdateTime.AsTime())
	}

	if len(rtPlatDef.PlatformDef.Clusters) <= 0 {
		return nil, fmt.Errorf("No clusters defined")
	}
	for idx, cluster := range rtPlatDef.PlatformDef.Clusters {
		if cluster.Name == "" {
			return nil, fmt.Errorf("Cluster %d missing name field", idx)
		}
		if cluster.Brokers == "" {
			return nil, fmt.Errorf("Cluster '%s' missing brokers field", cluster.Name)
		}
		if cluster.IsAdmin {
			if rtPlatDef.AdminCluster != "" {
				return nil, fmt.Errorf("More than one admin cluster")
			}
			rtPlatDef.AdminCluster = cluster.Name
		}
		// verify clusters only appear once
		if _, ok := rtPlatDef.Clusters[cluster.Name]; ok {
			return nil, fmt.Errorf("Cluster '%s' appears more than once in Platform '%s' definition", cluster.Name, rtPlatDef.PlatformDef.Name)
		}
		rtPlatDef.Clusters[cluster.Name] = cluster
	}
	if rtPlatDef.AdminCluster == "" {
		return nil, fmt.Errorf("No admin cluster defined")
	}

	if len(rtPlatDef.PlatformDef.StorageTargets) <= 0 {
		return nil, fmt.Errorf("No storage targets defined")
	}
	for idx, sttg := range rtPlatDef.PlatformDef.StorageTargets {
		if sttg.Name == "" {
			return nil, fmt.Errorf("Storage target %d missing name field", idx)
		}
		if sttg.Type == "" {
			return nil, fmt.Errorf("Storage target '%s' missing type field", sttg.Name)
		}
		if sttg.IsPrimary {
			if rtPlatDef.PrimaryStorageTarget != "" {
				return nil, fmt.Errorf("More than one primary storage target")
			}
			rtPlatDef.PrimaryStorageTarget = sttg.Name
		}
		// verify clusters only appear once
		if _, ok := rtPlatDef.StorageTargets[sttg.Name]; ok {
			return nil, fmt.Errorf("Storage target '%s' appears more than once in Platform '%s' definition", sttg.Name, rtPlatDef.PlatformDef.Name)
		}
		rtPlatDef.StorageTargets[sttg.Name] = sttg
	}
	if rtPlatDef.PrimaryStorageTarget == "" {
		return nil, fmt.Errorf("No primary storage target defined")
	}

	requiredTopics := map[Concern_Type][]string{
		Concern_GENERAL: {"admin", "error"},
		Concern_BATCH:   {"admin", "error"},
		Concern_APECS:   {"admin", "process", "error", "complete", "storage", "storage-scnd"},
	}

	for idx, concern := range rtPlatDef.PlatformDef.Concerns {
		if concern.Name == "" {
			return nil, fmt.Errorf("Concern %d missing name field", idx)
		}

		var topicNames []string
		// build list of topicNames for validation steps below
		for _, topics := range concern.Topics {
			topicNames = append(topicNames, topics.Name)
		}

		// validate our expected required topics are there, add any with defaults if not present
		for _, req := range requiredTopics[concern.Type] {
			if !contains(topicNames, req) {
				// conern.Topics will get initialized with reasonable defaults during topic validation below
				concern.Topics = append(concern.Topics, &Concern_Topics{Name: req})
			}
		}

		// ensure APECS concern only has required topics
		if concern.Type == Concern_APECS {
			// simple len check is adequate since we added all required above
			if len(requiredTopics[concern.Type]) != len(concern.Topics) {
				return nil, fmt.Errorf("ApecsConcern %d contains invalid command %+v required vs %+v total", idx, requiredTopics, concern.Topics)
			}
		}

		// validate all topics definitions
		for idx, _ := range concern.Topics {
			concern.Topics[idx] = initTopics(concern.Topics[idx], rtPlatDef.AdminCluster, concern.Type, rtPlatDef.PlatformDef.StorageTargets)
			if err := validateTopics(concern, concern.Topics[idx], rtPlatDef.Clusters); err != nil {
				return nil, fmt.Errorf("Concern '%s' has invalid '%s' Topics: %s", concern.Name, concern.Topics[idx].Name, err.Error())
			}
		}

		// verify concerns only appear once
		if _, ok := rtPlatDef.Concerns[concern.Name]; ok {
			return nil, fmt.Errorf("Concern '%s' appears more than once in Platform '%s' definition", concern.Name, rtPlatDef.PlatformDef.Name)
		}
		rtConc, err := newRtConcern(&rtPlatDef, concern)
		if err != nil {
			return nil, err
		}
		rtPlatDef.Concerns[concern.Name] = rtConc
	}

	return &rtPlatDef, nil
}

var singlePartitionTopics = map[string]bool{
	"admin":    true,
	"error":    true,
	"complete": true,
}

func validateTopics(concern *Concern, topics *Concern_Topics, clusters map[string]*Cluster) error {
	if topics.Name == "" {
		return fmt.Errorf("Topics missing Name field: %s", topics.Name)
	}

	if singlePartitionTopics[topics.Name] {
		if topics.Current != nil && topics.Current.PartitionCount != 1 {
			return fmt.Errorf("'%s' Topics must have exactly 1 partition", topics.Name)
		}
		if topics.Future != nil && topics.Future.PartitionCount != 1 {
			return fmt.Errorf("'%s' Topics must have exactly 1 partition", topics.Name)
		}
	}

	if topics.Current == nil {
		return fmt.Errorf("Topics missing Current Topic: %s", topics.Name)
	} else {
		if err := validateTopic(topics.Current, clusters); err != nil {
			return err
		}
	}
	if topics.Future != nil {
		if err := validateTopic(topics.Future, clusters); err != nil {
			return err
		}
		if topics.Current.Generation != topics.Future.Generation+1 {
			return fmt.Errorf("Future generation not Current + 1")
		}
	}

	for idx, consProg := range topics.ConsumerPrograms {
		if consProg != nil {
			if consProg.Name == "" {
				return fmt.Errorf("Program cannot have blank Name: %d", idx)
			}
			if consProg.Abbrev == "" {
				return fmt.Errorf("Program cannot have blank Abbrev: %s", consProg.Name)
			}
		}
	}
	return nil
}

func validateTopic(topic *Concern_Topic, clusters map[string]*Cluster) error {
	if topic.Generation == 0 {
		return fmt.Errorf("Topic missing Generation field")
	}
	if topic.Cluster == "" {
		return fmt.Errorf("Topic missing Cluster field")
	}
	if _, ok := clusters[topic.Cluster]; !ok {
		return fmt.Errorf("Topic refers to non-existent cluster: '%s'", topic.Cluster)
	}
	if topic.PartitionCount < 1 || topic.PartitionCount > MAX_PARTITION {
		return fmt.Errorf("Topic with out of bounds PartitionCount %d", topic.PartitionCount)
	}
	return nil
}

func uncommittedGroupName(topic string, partition int) string {
	return fmt.Sprintf("__%s_%d__non_comitted_group", topic, partition)
}

func uncommittedGroupNameAllPartitions(topic string) string {
	return fmt.Sprintf("__%s_ALL__non_comitted_group", topic)
}

func (plat *Platform) cobraPlatReplace(cmd *cobra.Command, args []string) {
	ctx, span := plat.telem.StartFunc(context.Background())
	defer span.End()

	slog := log.With().
		Str("Brokers", plat.settings.AdminBrokers).
		Str("PlatformPath", plat.settings.PlatformFilePath).
		Logger()

	// read platform conf file and deserialize
	conf, err := ioutil.ReadFile(plat.settings.PlatformFilePath)
	if err != nil {
		span.SetStatus(otel_codes.Error, err.Error())
		slog.Fatal().
			Err(err).
			Msg("Failed to ReadFile")
	}
	platDef := PlatformDef{}
	err = protojson.Unmarshal(conf, proto.Message(&platDef))
	if err != nil {
		span.SetStatus(otel_codes.Error, err.Error())
		slog.Fatal().
			Err(err).
			Msg("Failed to unmarshal platform")
	}

	platDef.UpdateTime = timestamppb.Now()

	// create an rtPlatformDef so we run the validations that involves
	rtPlatDef, err := newRtPlatformDef(&platDef, plat.name, plat.environment)
	if err != nil {
		span.SetStatus(otel_codes.Error, err.Error())
		slog.Fatal().
			Err(err).
			Msg("Failed to newRtPlatform")
	}
	jsonBytes, _ := protojson.Marshal(proto.Message(rtPlatDef.PlatformDef))
	log.Info().
		Str("PlatformJson", string(jsonBytes)).
		Msg("Platform parsed")

	// connect to kafka and make sure we have our platform topics
	err = createPlatformTopics(context.Background(), plat.settings.AdminBrokers, platDef.Name, platDef.Environment)
	if err != nil {
		span.SetStatus(otel_codes.Error, err.Error())
		slog.Fatal().
			Err(err).
			Str("Platform", platDef.Name).
			Msg("Failed to create platform topics")
	}

	platformTopic := PlatformTopic(platDef.Name, platDef.Environment)
	slog = slog.With().
		Str("Topic", platformTopic).
		Logger()

	// At this point we are guaranteed to have a platform admin topic
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	kafkaLogCh := make(chan kafka.LogEvent)
	go printKafkaLogs(ctx, kafkaLogCh)

	prod, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  plat.settings.AdminBrokers,
		"acks":               -1,     // acks required from all in-sync replicas
		"message.timeout.ms": 600000, // 10 minutes

		"go.logs.channel.enable": true,
		"go.logs.channel":        kafkaLogCh,
	})
	if err != nil {
		span.SetStatus(otel_codes.Error, err.Error())
		slog.Fatal().
			Err(err).
			Msg("Failed to NewProducer")
	}
	defer func() {
		log.Warn().
			Str("Brokers", plat.settings.AdminBrokers).
			Msg("Closing kafka producer")
		prod.Close()
		log.Warn().
			Str("Brokers", plat.settings.AdminBrokers).
			Msg("Closed kafka producer")
	}()

	msg, err := newKafkaMessage(&platformTopic, 0, &platDef, Directive_PLATFORM, ExtractTraceParent(ctx))
	if err != nil {
		span.SetStatus(otel_codes.Error, err.Error())
		slog.Fatal().
			Err(err).
			Msg("Failed to kafkaMessage")
	}

	produce := func() {
		err := prod.Produce(msg, nil)
		if err != nil {
			span.SetStatus(otel_codes.Error, err.Error())
			slog.Fatal().
				Err(err).
				Msg("Failed to Produce")
		}
	}

	produce()

	// check channel for delivery event
	timer := time.NewTimer(10 * time.Second)
Loop:
	for {
		select {
		case <-timer.C:
			slog.Fatal().
				Msg("Timeout producing platform message")
		case ev := <-prod.Events():
			msgEv, ok := ev.(*kafka.Message)
			if !ok {
				slog.Warn().
					Msg("Non *kafka.Message event received from producer")
			} else {
				if msgEv.TopicPartition.Error != nil {
					slog.Warn().
						Err(msgEv.TopicPartition.Error).
						Msg("Error reported while producing platform message, trying again after a delay")
					time.Sleep(1 * time.Second)
					produce()
				} else {
					slog.Info().
						Msg("Platform config successfully produced")
					break Loop
				}
			}
		}
	}
}
