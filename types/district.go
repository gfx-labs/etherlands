package types

import (
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

	playerPermissions PlayerPermissionMap
	groupPermissions  GroupPermissionMap

	mutex sync.RWMutex

	key   FamilyKey
	world *World
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

func (D *District) PlayerPermissions() PlayerPermissionMap {
	D.mutex.RLock()
	defer D.mutex.RUnlock()
	return D.playerPermissions
}

func (D *District) GroupPermissions() GroupPermissionMap {
	D.mutex.RLock()
	defer D.mutex.RUnlock()
	return D.groupPermissions
}

func BuildDistrictGroupPermissionVector(
	builder *flatbuffers.Builder,
	target GroupPermissionMap,
) flatbuffers.UOffsetT {
	gp_o := BuildGroupPermissions(builder, target)
	proto.DistrictStartGroupPermissionsVector(builder, len(gp_o))
	for _, v := range gp_o {
		builder.PrependUOffsetT(v)
	}
	return builder.EndVector(len(gp_o))
}

func BuildDistrictPlayerPermissionVector(
	builder *flatbuffers.Builder,
	target PlayerPermissionMap,
) flatbuffers.UOffsetT {
	pp_o := BuildPlayerPermissions(builder, target)
	proto.DistrictStartPlayerPermissionsVector(builder, len(pp_o))
	for _, v := range pp_o {
		builder.PrependUOffsetT(v)
	}
	return builder.EndVector(len(pp_o))
}

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

	player_permission_offset := BuildDistrictPlayerPermissionVector(builder, D.PlayerPermissions())
	group_permission_offset := BuildDistrictGroupPermissionVector(builder, D.GroupPermissions())

	proto.DistrictStart(builder)

	proto.DistrictAddChainId(builder, D.DistrictId())

	proto.DistrictAddNickname(builder, nickname_offset)

	if D.Owner() != nil {
		owner_uuid_offset := BuildUUID(builder, D.Owner().MinecraftId())
		proto.DistrictAddOwnerUuid(builder, owner_uuid_offset)
	}
	proto.DistrictAddOwnerAddress(builder, owner_address_offset)
	proto.DistrictAddPlots(builder, plots_offset)
	proto.DistrictAddGroupPermissions(builder, group_permission_offset)
	proto.DistrictAddPlayerPermissions(builder, player_permission_offset)

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
