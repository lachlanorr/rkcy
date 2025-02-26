// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package mgmt

import (
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"

	"github.com/lachlanorr/rocketcycle/pkg/rkcy"
	"github.com/lachlanorr/rocketcycle/pkg/rkcypb"
)

type ConcernAdminMessage struct {
	Directive             rkcypb.Directive
	Timestamp             time.Time
	Offset                int64
	ConcernAdminDirective *rkcypb.ConcernAdminDirective
}

func ConsumeConcernAdminTopic(
	ctx context.Context,
	wg *sync.WaitGroup,
	strmprov rkcy.StreamProvider,
	platform string,
	environment string,
	adminBrokers string,
	ch chan<- *ConcernAdminMessage,
	concern string,
	readyCh chan<- bool,
) {
	wg.Add(1)
	go ConsumeACETopic(
		ctx,
		wg,
		strmprov,
		platform,
		environment,
		adminBrokers,
		concern,
		rkcy.ADMIN,
		rkcypb.Directive_CONCERN_ADMIN,
		PAST_LAST_MATCH,
		func(rawMsg *RawMessage) {
			cncAdminDir := &rkcypb.ConcernAdminDirective{}
			err := proto.Unmarshal(rawMsg.Value, cncAdminDir)
			if err != nil {
				log.Error().
					Err(err).
					Msg("Failed to Unmarshal ConcernAdminDirective")
				return
			}

			ch <- &ConcernAdminMessage{
				Directive:             rawMsg.Directive,
				Timestamp:             rawMsg.Timestamp,
				Offset:                rawMsg.Offset,
				ConcernAdminDirective: cncAdminDir,
			}
		},
		readyCh,
	)
}
