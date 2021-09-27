package types

import (
	"errors"
	"fmt"

	proto "github.com/gfx-labs/etherlands/proto"
	flatbuffers "github.com/google/flatbuffers/go"
	uuid "github.com/google/uuid"
)

func (W *World) TownCount() int {
	return len(W.towns)
}

func (W *World) Towns() []*Town {
	output := []*Town{}
	W.towns_lock.RLock()
	defer W.towns_lock.RUnlock()
	for _, v := range W.towns {
		if v != nil {
			output = append(output, v)
		}
	}
	return output
}

func (W *World) GetTown(name string) (*Town, error) {
	W.towns_lock.RLock()
	if val, ok := W.towns[NewTownKey(name)]; ok {
		W.towns_lock.RUnlock()
		return val, nil
	}
	W.towns_lock.RUnlock()
	//obtain a write lock
	W.towns_lock.Lock()
	defer W.towns_lock.Unlock()
	// if not in live cache, see if town file exists
	if res, err := W.LoadTown(name); err == nil {
		return res, nil
	}
	return nil, errors.New(fmt.Sprintf("town %s could not be found", name))
}

type town_create_input struct {
	name  string
	owner *Gamer
}

func (W *World) CreateTown(name string, owner *Gamer) error {
	town, err := W.GetTown(name)
	if err != nil && town == nil {
		newTown := W.initTown(name)
		newTown.SetOwner(owner.MinecraftId())
		newTown.Members()
		W.UpdateTown(newTown)
		return nil
	}
	return errors.New(fmt.Sprintf("town with name %s already exists", name))
}

func (W *World) initTown(name string) *Town {
	return &Town{
		W:                        W,
		members:                  make(map[uuid.UUID]struct{}),
		districts:                make([]uint64, 0),
		groups:                   make(map[string]*Group),
		defaultPlayerPermissions: NewPlayerPermissionMap(),
		defaultGroupPermissions:  NewGroupPermissionMap(),
		district_player_lock:     NewDistrictLock(),
		district_group_lock:      NewDistrictLock(),
	}

}

func (W *World) LoadTown(name string) (*Town, error) {
	bytes, err := ReadStruct("towns", name)
	if err != nil {
		return nil, err
	}
	if len(bytes) < 8 {
		return nil, errors.New(fmt.Sprintf("Empty file for %s", name))
	}
	read_town := proto.GetRootAsTown(bytes, 0)
	pending_town := W.initTown(name)

	pending_town.districts = make([]uint64, read_town.DistrictsLength())
	pending_town.name = string(read_town.Name())
	pending_town.owner = ProtoResolveUUID(read_town.Owner(nil))
	for i := 0; i < read_town.MembersLength(); i++ {
		var puuid proto.UUID
		read_town.Members(&puuid, i)
		pending_town.AddMember(ProtoResolveUUID(&puuid))
	}
	for i := 0; i < read_town.DistrictsLength(); i++ {
		pending_town.districts[i] = read_town.Districts(i)
	}
	for i := 0; i < read_town.GroupsLength(); i++ {
		var group proto.Group
		read_town.Groups(&group, i)
		uuids := make([]uuid.UUID, group.MembersLength())
		for j := 0; j < group.MembersLength(); j++ {
			puuid := proto.UUID{}
			group.Members(&puuid, j)
			uuids[j] = ProtoResolveUUID(&puuid)
		}
		pending_town.groups[string(group.Name())] =
			&Group{
				name:    string(group.Name()),
				members: uuids,
			}
	}
	default_group_map := read_town.DefaultGroupPermissions(nil)
	for i := 0; i < default_group_map.PermissionsLength(); i++ {
		var perm proto.GroupPermission
		default_group_map.Permissions(&perm, i)
		pending_town.DefaultGroupPermissions().insert(
			string(perm.Group()),
			perm.Flag(),
			perm.Value(),
		)
	}
	default_player_map := read_town.DefaultPlayerPermissions(nil)
	for i := 0; i < default_player_map.PermissionsLength(); i++ {
		var perm proto.PlayerPermission
		default_player_map.Permissions(&perm, i)
		pending_town.DefaultPlayerPermissions().insert(
			ProtoResolveUUID(perm.MinecraftId(nil)),
			perm.Flag(),
			perm.Value(),
		)
	}
	district_group_maps := read_town.DistrictGroupPermissions(nil)
	for h := 0; h < district_group_maps.DistrictsLength(); h++ {
		district_group_map := proto.GroupPermissionMap{}
		district_group_maps.Permissions(&district_group_map, h)
		for i := 0; i < district_group_map.PermissionsLength(); i++ {
			var perm proto.GroupPermission
			district_group_map.Permissions(&perm, i)
			pending_town.DistrictGroupPermission(district_group_maps.Districts(h)).insert(
				string(perm.Group()),
				perm.Flag(),
				perm.Value(),
			)
		}
	}

	district_player_maps := read_town.DistrictPlayerPermissions(nil)
	for h := 0; h < district_player_maps.DistrictsLength(); h++ {
		district_player_map := proto.PlayerPermissionMap{}
		district_player_maps.Permissions(&district_player_map, h)
		for i := 0; i < district_player_map.PermissionsLength(); i++ {
			var perm proto.PlayerPermission
			district_player_map.Permissions(&perm, i)
			pending_town.DistrictPlayerPermission(district_player_maps.Districts(h)).insert(
				ProtoResolveUUID(perm.MinecraftId(nil)),
				perm.Flag(),
				perm.Value(),
			)
		}
	}

	//W.UpdateTown(pending_town)
	return pending_town, nil
}

func (T *Town) Save() error {
	builder := flatbuffers.NewBuilder(1024)
	// create default player permission map
	player_permission_offset := BuildPlayerPermissionMap(builder, T.DefaultPlayerPermissions())
	// create default group permission map
	group_permission_offset := BuildGroupPermissionMap(builder, T.DefaultGroupPermissions())

	// create district player permission map
	district_player_permission_offset := BuildDistrictPlayerPermissionMap(
		builder,
		T.districtPlayerPermissions,
	)
	// create district group permission map
	district_group_permission_offset := BuildDistrictGroupPermissionMap(
		builder,
		T.districtGroupPermissions,
	)

	// create districts vector
	proto.TownStartDistrictsVector(builder, len(T.Districts()))
	for _, v := range T.Districts() {
		builder.PrependUint64(v)
	}
	districts_offset := builder.EndVector(len(T.Districts()))

	// prepare town member vector
	town_members := T.Members()
	me_o := make([]flatbuffers.UOffsetT, len(town_members))
	idx := 0
	for k := range town_members {
		me_o[idx] = BuildUUID(builder, k)
		idx = idx + 1
	}

	// create town member vector
	proto.TownStartMembersVector(builder, len(town_members))
	for _, v := range me_o {
		builder.PrependUOffsetT(v)
	}
	member_vector := builder.EndVector(len(town_members))

	town_name := builder.CreateString(T.Name())

	//owner_id := BuildUUID(builder, T.Owner())

	//create town table
	proto.TownStart(builder)
	//town name
	proto.TownAddName(builder, town_name)
	//owner
	//proto.TownAddOwner(builder, owner_id)
	//members
	proto.TownAddMembers(builder, member_vector)

	//districts
	proto.TownAddDistricts(builder, districts_offset)
	//perms
	proto.TownAddDefaultGroupPermissions(builder, group_permission_offset)
	proto.TownAddDefaultPlayerPermissions(builder, player_permission_offset)
	proto.TownAddDistrictPlayerPermissions(builder, district_player_permission_offset)
	proto.TownAddDistrictGroupPermissions(builder, district_group_permission_offset)

	//finish
	town_offset := proto.TownEnd(builder)
	builder.Finish(town_offset)
	buf := builder.FinishedBytes()

	return WriteStruct("town", T.Name(), buf)
}

func BuildGroupVector(
	builder *flatbuffers.Builder,
	target map[string]*Group,
) flatbuffers.UOffsetT {
	go_a := make([]flatbuffers.UOffsetT, 0)
	for k, v := range target {
		name := builder.CreateString(k)
		memes := v.Members()
		gmo := make([]flatbuffers.UOffsetT, len(memes))
		for j := 0; j < len(memes); j++ {
			gmo[j] = BuildUUID(builder, memes[j])
		}
		proto.GroupStart(builder)
		for j := 0; j < len(gmo); j++ {
			proto.GroupAddMembers(builder, gmo[j])
		}
		proto.GroupAddName(builder, name)
		go_a = append(go_a, proto.GroupEnd(builder))
	}
	proto.TownStartGroupsVector(builder, len(go_a))
	for i := 0; i < len(go_a); i++ {
		builder.PrependUOffsetT(go_a[i])
	}
	return builder.EndVector(len(go_a))

}

func BuildDistrictGroupPermissionMap(
	builder *flatbuffers.Builder,
	target map[uint64]*GroupPermissionMap,
) flatbuffers.UOffsetT {

	district_ids := []uint64{}
	group_perms := []flatbuffers.UOffsetT{}
	map_count := 0
	for k, v := range target {
		map_count = map_count + 1
		district_ids = append(district_ids, k)
		group_perms = append(group_perms, BuildGroupPermissionMap(builder, v))
	}
	proto.DistrictGroupPermissionMapStartDistrictsVector(builder, map_count)
	for _, v := range district_ids {
		builder.PrependUint64(v)
	}
	districts := builder.EndVector(map_count)
	proto.DistrictGroupPermissionMapStartPermissionsVector(builder, map_count)
	for _, v := range group_perms {
		builder.PrependUOffsetT(v)
	}
	permissions := builder.EndVector(map_count)
	proto.DistrictGroupPermissionMapStart(builder)
	proto.DistrictGroupPermissionMapAddDistricts(builder, districts)
	proto.DistrictGroupPermissionMapAddPermissions(builder, permissions)
	return proto.DistrictGroupPermissionMapEnd(builder)
}

func BuildDistrictPlayerPermissionMap(
	builder *flatbuffers.Builder,
	target map[uint64]*PlayerPermissionMap,
) flatbuffers.UOffsetT {
	district_ids := []uint64{}
	gamer_perms := []flatbuffers.UOffsetT{}
	map_count := 0
	for k, v := range target {
		map_count = map_count + 1
		district_ids = append(district_ids, k)
		gamer_perms = append(gamer_perms, BuildPlayerPermissionMap(builder, v))
	}
	proto.DistrictPlayerPermissionMapStartDistrictsVector(builder, map_count)
	for _, v := range district_ids {
		builder.PrependUint64(v)
	}
	districts := builder.EndVector(map_count)
	proto.DistrictGroupPermissionMapStartPermissionsVector(builder, map_count)
	for _, v := range gamer_perms {
		builder.PrependUOffsetT(v)
	}
	permissions := builder.EndVector(map_count)
	proto.DistrictPlayerPermissionMapStart(builder)
	proto.DistrictPlayerPermissionMapAddDistricts(builder, districts)
	proto.DistrictPlayerPermissionMapAddPermissions(builder, permissions)
	return proto.DistrictPlayerPermissionMapEnd(builder)
}

func BuildGroupPermissionMap(
	builder *flatbuffers.Builder,
	target *GroupPermissionMap,
) flatbuffers.UOffsetT {
	gp_o := BuildGroupPermissions(builder, target)
	proto.GroupPermissionMapStartPermissionsVector(builder, len(gp_o))
	for _, v := range gp_o {
		builder.PrependUOffsetT(v)
	}
	return builder.EndVector(len(gp_o))
}

func BuildPlayerPermissionMap(
	builder *flatbuffers.Builder,
	target *PlayerPermissionMap,
) flatbuffers.UOffsetT {
	pp_o := BuildPlayerPermissions(builder, target)
	proto.PlayerPermissionMapStartPermissionsVector(builder, len(pp_o))
	for _, v := range pp_o {
		builder.PrependUOffsetT(v)
	}
	return builder.EndVector(len(pp_o))
}
