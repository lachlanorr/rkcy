// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rckafka

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"google.golang.org/protobuf/proto"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func FindHeader(msg *kafka.Message, key string) []byte {
	for _, hdr := range msg.Headers {
		if key == hdr.Key {
			return hdr.Value
		}
	}
	return nil
}

func AdminTopic(internalName string) string {
	return fmt.Sprintf("rc.%s.admin", internalName)
}

func CreateAdminTopic(ctx context.Context, bootstrapServers string, internalName string) (string, error) {
	topicName := AdminTopic(internalName)

	// connect to kafka and make sure we have our platform topic
	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
	})
	if err != nil {
		return "", errors.New("Failed to NewAdminClient")
	}

	md, err := admin.GetMetadata(nil, true, 1000)
	if err != nil {
		return "", errors.New("Failed to GetMetadata")
	}

	_, ok := md.Topics[topicName]
	if !ok { // platform topic doesn't exist
		result, err := admin.CreateTopics(
			context.Background(),
			[]kafka.TopicSpecification{
				{
					Topic:             topicName,
					NumPartitions:     1,
					ReplicationFactor: len(md.Brokers),
					Config: map[string]string{
						"retention.ms":    "-1",
						"retention.bytes": strconv.Itoa(10 * 1024 * 1024),
					},
				},
			},
			nil,
		)
		if err != nil {
			return "", errors.New("Failed to create metadata topic")
		}
		for _, res := range result {
			if res.Error.Code() != kafka.ErrNoError {
				return "", errors.New("Failed to create metadata topic")
			}
		}
	}
	return topicName, nil
}

func MsgTypeName(msg proto.Message) string {
	msgR := msg.ProtoReflect()
	desc := msgR.Descriptor()
	return string(desc.FullName())
}

func MsgTypeHeaders(msg proto.Message) []kafka.Header {
	return []kafka.Header{{Key: "type", Value: []byte(MsgTypeName(msg))}}
}
