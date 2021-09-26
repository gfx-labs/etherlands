package types

import (
	"encoding/binary"

	proto "github.com/gfx-labs/etherlands/proto"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/google/uuid"
)

func BreakUUID(id uuid.UUID) (uint64, uint64) {
	l1 := binary.BigEndian.Uint64(id[:8])
	l2 := binary.BigEndian.Uint64(id[8:])
	return l1, l2
}

func (T *Town) Save() error {
	builder := flatbuffers.NewBuilder(1024)
	// create default player permission vector
	player_permission_offset := BuildTownPlayerPermissionVector(builder, T.defaultPlayerPermissions)
	// create default group permission vector
	group_permission_offset := BuildTownGroupPermissionVector(builder, T.defaultGroupPermissions)

	// create districts vector
	proto.TownStartDistrictsVector(builder, len(T.Districts()))
	for _, v := range T.Districts() {
		builder.PrependUint64(v.DistrictId())
	}
	districts_offset := builder.EndVector(len(T.Districts()))

	// create town manager vector
	town_managers := T.Managers()
	proto.TownStartManagersVector(builder, len(town_managers))
	for _, v := range town_managers {
		manager_offset := BuildUUID(builder, v.MinecraftId())
		builder.PrependUOffsetT(manager_offset)
	}
	manager_vector := builder.EndVector(len(town_managers))

	// create town member vector
	town_members := T.Members()
	proto.TownStartMembersVector(builder, len(town_members))
	for _, v := range town_members {
		member_offset := BuildUUID(builder, v.MinecraftId())
		builder.PrependUOffsetT(member_offset)
	}
	member_vector := builder.EndVector(len(town_members))

	//create town table
	proto.TownStart(builder)
	//town name
	town_name := builder.CreateString(T.Name())
	proto.TownAddName(builder, town_name)

	//owner
	owner_id := BuildUUID(builder, T.Owner().MinecraftId())
	proto.TownAddOwner(builder, owner_id)

	//members
	proto.TownAddMembers(builder, member_vector)

	//managers
	proto.TownAddManagers(builder, manager_vector)
	//districts
	proto.TownAddDistricts(builder, districts_offset)
	//perms
	proto.TownAddDefaultGroupPermissions(builder, group_permission_offset)
	proto.TownAddDefaultPlayerPermissions(builder, player_permission_offset)

	//finish
	town_offset := proto.TownEnd(builder)
	builder.Finish(town_offset)
	buf := builder.FinishedBytes()

	return WriteStruct("town", T.Name(), buf)
}
