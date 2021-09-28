package types

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	proto "github.com/gfx-labs/etherlands/proto"
	flatbuffers "github.com/google/flatbuffers/go"
)

type District struct {
	district_id   uint64
	owner         *Gamer
	owner_address string

	nickname *[24]byte

	mutex sync.RWMutex

	key FamilyKey
	W   *World

	town      string
	town_lock sync.RWMutex
}

func (D *District) DelegateTown(gamer *Gamer) error {
	if gamer.HasTown() {
		if gamer.Address() == D.OwnerAddress() {
			D.town_lock.Lock()
			D.town = gamer.Town()
			D.town_lock.Unlock()
			D.Update()
			return nil
		}
		return errors.New(fmt.Sprintf("You must own district %s to delegate it", D.StringName()))
	}
	return errors.New("You must be in a town to delegate")
}

func (D *District) Reclaim(gamer *Gamer) error {
	if gamer.Address() == D.OwnerAddress() {
		D.town_lock.Lock()
		D.town = ""
		D.town_lock.Unlock()
		D.Update()
		return nil
	}
	return errors.New(fmt.Sprintf("You must own district %s to reclaim it", D.StringName()))
}

func (D *District) HasTown() bool {
	D.town_lock.RLock()
	defer D.town_lock.RUnlock()
	return D.town != ""
}

func (D *District) Town() string {
	D.town_lock.RLock()
	defer D.town_lock.RUnlock()
	return D.town
}

func (D *District) DistrictId() uint64 {
	return D.district_id
}

func (D *District) SetNickname(newName [24]byte) {
	D.mutex.Lock()
	defer D.mutex.Unlock()
	D.nickname = &newName
}

func (D *District) Nickname() *[24]byte {
	D.mutex.RLock()
	defer D.mutex.RUnlock()
	return D.nickname
}

func (D *District) StringName() string {
	pending := Parse24Name(*D.Nickname())
	if pending == "" {
		return fmt.Sprintf("#%d", D.district_id)
	}
	return pending
}

func (D *District) OwnerAddress() string {
	D.mutex.RLock()
	defer D.mutex.RUnlock()
	return D.owner_address
}

func (D *District) SetOwnerAddress(addr string) {
	D.mutex.Lock()
	defer D.mutex.Unlock()
	D.owner_address = addr
}

func (D *District) Owner() *Gamer {
	D.mutex.RLock()
	defer D.mutex.RUnlock()
	return D.owner
}

//func (D *District) PlayerPermissions() PlayerPermissionMap {
//	D.mutex.RLock()
//	defer D.mutex.RUnlock()
//	return D.playerPermissions
//}
//
//func (D *District) GroupPermissions() GroupPermissionMap {
//	D.mutex.RLock()
//	defer D.mutex.RUnlock()
//	return D.groupPermissions
//}

func (D *District) GetKey() FamilyKey {
	return D.key
}

func (D *District) Save() error {
	builder := flatbuffers.NewBuilder(1024)

	nickname_offset := builder.CreateByteVector((D.Nickname())[:])
	proto.DistrictStartPlotsVector(builder, len(D.Plots()))
	for _, v := range D.Plots() {
		builder.PrependUint64(v.PlotId())
	}
	plots_offset := builder.EndVector(len(D.Plots()))
	owner_address_offset := builder.CreateString(D.OwnerAddress())

	var owner_uuid_offset flatbuffers.UOffsetT
	if D.Owner() != nil {
		owner_uuid_offset = BuildUUID(builder, D.Owner().MinecraftId())
	}

	proto.DistrictStart(builder)

	proto.DistrictAddChainId(builder, D.DistrictId())

	proto.DistrictAddNickname(builder, nickname_offset)
	if D.Owner() != nil {
		proto.DistrictAddOwnerUuid(builder, owner_uuid_offset)
	}
	proto.DistrictAddOwnerAddress(builder, owner_address_offset)
	proto.DistrictAddPlots(builder, plots_offset)

	//finish
	district_offset := proto.DistrictEnd(builder)
	builder.Finish(district_offset)
	buf := builder.FinishedBytes()

	return WriteStruct("districts", strconv.FormatUint(D.DistrictId(), 10), buf)
}

func NewDistrictKey(district_id uint64) FamilyKey {
	return FamilyKey{
		datatype: DISTRICT_FAMILY,
		subkey:   strconv.FormatUint(district_id, 10),
	}
}
