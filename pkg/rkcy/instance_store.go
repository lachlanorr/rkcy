// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package rkcy

import (
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/lachlanorr/rocketcycle/pkg/rkcypb"
)

type InstanceRecord struct {
	Instance   []byte
	Related    []byte
	CmpdOffset *rkcypb.CompoundOffset
	LastAccess time.Time
}

type InstanceStore struct {
	instances map[string]*InstanceRecord
}

func NewInstanceStore() *InstanceStore {
	instStore := InstanceStore{
		instances: make(map[string]*InstanceRecord),
	}
	return &instStore
}

func (instStore *InstanceStore) Get(key string) *InstanceRecord {
	rec, ok := instStore.instances[key]
	if ok {
		rec.LastAccess = time.Now()
		return rec
	}
	return nil
}

func (instStore *InstanceStore) GetPacked(key string) []byte {
	rec, ok := instStore.instances[key]
	if ok {
		rec.LastAccess = time.Now()
		return PackPayloads(rec.Instance, rec.Related)
	}
	return nil
}

func (instStore *InstanceStore) GetInstance(key string) []byte {
	rec := instStore.Get(key)
	if rec != nil {
		return rec.Instance
	}
	return nil
}

func (instStore *InstanceStore) GetRelated(key string) []byte {
	rec := instStore.Get(key)
	if rec != nil {
		return rec.Related
	}
	return nil
}

func (instStore *InstanceStore) Set(key string, instance []byte, related []byte, cmpdOffset *rkcypb.CompoundOffset) {
	rec, ok := instStore.instances[key]
	if ok {
		if !OffsetGT(cmpdOffset, rec.CmpdOffset) {
			log.Error().
				Str("Key", key).
				Msgf("Out of order InstanceStore.Set: new (%+v), old (%+v)", cmpdOffset, rec.CmpdOffset)
			return
		}
		rec.Instance = instance
		rec.Related = related
		rec.CmpdOffset = cmpdOffset
		rec.LastAccess = time.Now()
	} else {
		instStore.instances[key] = &InstanceRecord{Instance: instance, Related: related, CmpdOffset: cmpdOffset, LastAccess: time.Now()}
	}
}

func (instStore *InstanceStore) SetInstance(key string, instance []byte, cmpdOffset *rkcypb.CompoundOffset) {
	rec, ok := instStore.instances[key]
	if ok {
		if !OffsetGT(cmpdOffset, rec.CmpdOffset) {
			log.Error().
				Str("Key", key).
				Msgf("Out of order InstanceStore.Set: new (%+v), old (%+v)", cmpdOffset, rec.CmpdOffset)
			return
		}
		rec.Instance = instance
		rec.CmpdOffset = cmpdOffset
		rec.LastAccess = time.Now()
	} else {
		instStore.instances[key] = &InstanceRecord{Instance: instance, CmpdOffset: cmpdOffset, LastAccess: time.Now()}
	}
}

func (instStore *InstanceStore) SetRelated(key string, related []byte, cmpdOffset *rkcypb.CompoundOffset) error {
	rec, ok := instStore.instances[key]
	if ok {
		if !OffsetGT(cmpdOffset, rec.CmpdOffset) {
			log.Error().
				Str("Key", key).
				Msgf("Out of order InstanceStore.Set: new (%+v), old (%+v)", cmpdOffset, rec.CmpdOffset)
			return nil
		}
		rec.Related = related
		rec.CmpdOffset = cmpdOffset
		rec.LastAccess = time.Now()
		return nil
	}
	log.Error().
		Str("Key", key).
		Msgf("Unable to set related on missing key")
	return fmt.Errorf("Unable to set related on missing key: %s", key)
}

func (instStore *InstanceStore) Remove(key string) {
	delete(instStore.instances, key)
}
