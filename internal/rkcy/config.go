// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rkcy

import (
	"context"
	"io/ioutil"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	admin_pb "github.com/lachlanorr/rocketcycle/build/proto/admin"
)

func confCommand(cmd *cobra.Command, args []string) {
	slog := log.With().
		Str("BootstrapServers", flags.bootstrapServers).
		Str("ConfigPath", flags.configFilePath).
		Logger()

	// read platform conf file and deserialize
	conf, err := ioutil.ReadFile(flags.configFilePath)
	if err != nil {
		slog.Fatal().
			Err(err).
			Msg("Failed to ReadFile")
	}
	plat := admin_pb.Platform{}
	err = protojson.Unmarshal(conf, proto.Message(&plat))
	if err != nil {
		slog.Fatal().
			Err(err).
			Msg("Failed to unmarshal platform")
	}
	platMar, err := proto.Marshal(&plat)
	if err != nil {
		slog.Fatal().
			Err(err).
			Msg("Failed to Marshal platform")
	}

	// connect to kafka and make sure we have our platform topic
	adminTopic, err := createAdminTopic(context.Background(), flags.bootstrapServers, plat.Name)
	if err != nil {
		slog.Fatal().
			Err(err).
			Msgf("Failed to createAdminTopic for platform %s", plat.Name)
	}
	slog = slog.With().
		Str("Topic", adminTopic).
		Logger()
	slog.Info().
		Msgf("Created platform admin topic: %s", adminTopic)

	// At this point we are guaranteed to have a platform admin topic
	prod, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": flags.bootstrapServers})
	if err != nil {
		slog.Fatal().
			Err(err).
			Msg("Failed to NewProducer")
	}
	defer prod.Close()

	err = prod.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &adminTopic, Partition: 0},
		Value:          platMar,
		Headers:        msgTypeHeaders(proto.Message(&plat)),
	}, nil)
	if err != nil {
		slog.Fatal().
			Err(err).
			Msg("Failed to Produce")
	}

	// check channel for delivery event
	timer := time.NewTimer(5 * time.Second)
	select {
	case <-timer.C:
		slog.Fatal().
			Msg("Timeout waiting for producer callback")
	case ev := <-prod.Events():
		msgEv, ok := ev.(*kafka.Message)
		if !ok {
			slog.Fatal().
				Msg("Non *kafka.Message event received from producer")
		} else {
			if msgEv.TopicPartition.Error != nil {
				slog.Fatal().
					Err(msgEv.TopicPartition.Error).
					Msg("Error reported while producing platform message")
			} else {
				slog.Info().
					Msg("Platform config successfully produced")
			}
		}
	}
}
