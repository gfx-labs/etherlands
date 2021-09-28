package types

import (
	"errors"
	"fmt"
	"time"

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
	// if not in live cache, see if town file exists
	if res, err := W.LoadTown(name); err == nil {
		return res, nil
	}
	return nil, errors.New(fmt.Sprintf("town [town.%s] could not be found", name))
}

type town_create_input struct {
	name  string
	owner *Gamer
}

func (W *World) CreateTown(name string, owner *Gamer) error {
	if owner.HasTown() {
		return errors.New("You must leave your town before creating one")
	}
	town, err := W.GetTown(name)
	if err != nil && town == nil {
		newTown := W.initTown(name)
		newTown.owner = owner.MinecraftId()
		owner.SetTown(name)
		owner.Update()
		W.UpdateTown(newTown)
		return nil
	}
	return errors.New(fmt.Sprintf("Town [town.%s] already exists", name))
}

func (W *World) initTown(name string) *Town {
	town := &Town{
		W:                         W,
		name:                      name,
		districts:                 make([]uint64, 0),
		teams:                     make(map[string]*Team),
		invites:                   make(map[uuid.UUID]time.Time),
		districtPlayerPermissions: NewDistrictPlayerPermissionMap(),
		districtTeamPermissions:   NewDistrictTeamPermissionMap(),
		district_player_lock:      NewMapLock(),
		district_team_lock:        NewMapLock(),
	}
	town.teams["manager"] = &Team{
		name:     "manager",
		priority: 100,
		members:  make(map[uuid.UUID]struct{}),
	}
	town.teams["member"] = &Team{name: "member", priority: -100}
	town.teams["outsider"] = &Team{name: "outsider", priority: -1}
	go town.ProcessInvites(15 * time.Minute)
	return town
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
	for i := 0; i < read_town.DistrictsLength(); i++ {
		pending_town.districts[i] = read_town.Districts(i)
	}
	for i := 0; i < read_town.TeamsLength(); i++ {
		var team proto.Team
		read_town.Teams(&team, i)
		team_members := make(map[uuid.UUID]struct{})
		for j := 0; j < team.MembersLength(); j++ {
			var puuid proto.UUID
			team.Members(&puuid, j)
			new_uuid := ProtoResolveUUID(&puuid)
			team_members[new_uuid] = struct{}{}
		}

		pending_town.teams[string(team.Name())] =
			&Team{
				name:     string(team.Name()),
				priority: team.Priority(),
				members:  team_members,
			}
	}
	district_team_maps := read_town.DistrictTeamPermissions(nil)
	if district_team_maps != nil {
		for h := 0; h < district_team_maps.DistrictsLength(); h++ {
			district_team_map := proto.TeamPermissionMap{}
			district_team_maps.Permissions(&district_team_map, h)
			for i := 0; i < district_team_map.PermissionsLength(); i++ {
				var perm proto.TeamPermission
				district_team_map.Permissions(&perm, i)
				pending_town.DistrictTeamPermissions().Insert(
					district_team_maps.Districts(h),
					string(perm.Team()),
					perm.Flag(),
					perm.Value(),
				)
			}
		}
	}

	district_player_maps := read_town.DistrictPlayerPermissions(nil)
	if district_player_maps != nil {
		for h := 0; h < district_player_maps.DistrictsLength(); h++ {
			district_player_map := proto.PlayerPermissionMap{}
			district_player_maps.Permissions(&district_player_map, h)
			for i := 0; i < district_player_map.PermissionsLength(); i++ {
				var perm proto.PlayerPermission
				district_player_map.Permissions(&perm, i)
				pending_town.DistrictPlayerPermissions().Insert(
					district_player_maps.Districts(h),
					ProtoResolveUUID(perm.MinecraftId(nil)),
					perm.Flag(),
					perm.Value(),
				)
			}
		}
	}

	go W.UpdateTown(pending_town)
	return pending_town, nil
}

func (T *Town) Save() error {
	builder := flatbuffers.NewBuilder(1024)
	// team vector
	team_vector := buildTeamVector(builder, T.Teams())
	// permission vectors
	T.districtPlayerPermissions.global.Lock()
	district_player_permission_offset := buildDistrictPlayerPermissionMap(
		builder,
		T.districtPlayerPermissions,
	)
	T.districtPlayerPermissions.global.Unlock()
	T.districtTeamPermissions.global.Lock()
	district_team_permission_offset := buildDistrictTeamPermissionMap(
		builder,
		T.districtTeamPermissions,
	)
	T.districtTeamPermissions.global.Unlock()

	// prepare town member vector
	town_members := T.Members()
	me_o := make([]flatbuffers.UOffsetT, len(town_members))
	idx := 0
	for k := range town_members {
		me_o[idx] = BuildUUID(builder, k)
		idx = idx + 1
	}

	town_name := builder.CreateString(T.Name())

	owner_id := BuildUUID(builder, T.Owner())

	//create town table
	proto.TownStart(builder)
	//owner
	proto.TownAddOwner(builder, owner_id)
	//town name
	proto.TownAddName(builder, town_name)

	//teams
	proto.TownAddTeams(builder, team_vector)

	//perms
	proto.TownAddDistrictPlayerPermissions(builder, district_player_permission_offset)
	proto.TownAddDistrictTeamPermissions(builder, district_team_permission_offset)

	//finish
	town_offset := proto.TownEnd(builder)
	builder.Finish(town_offset)
	buf := builder.FinishedBytes()

	return WriteStruct("towns", T.Name(), buf)
}

func buildTeamVector(
	builder *flatbuffers.Builder,
	target map[string]*Team,
) flatbuffers.UOffsetT {
	go_a := []flatbuffers.UOffsetT{}
	for k, v := range target {
		name := builder.CreateString(k)
		memes := v.Members()
		proto.TeamStartMembersVector(builder, len(memes))
		for k := range memes {
			builder.PrependUOffsetT(BuildUUID(builder, k))
		}
		members_vector := builder.EndVector(len(memes))
		proto.TeamStart(builder)
		proto.TeamAddMembers(builder, members_vector)
		proto.TeamAddPriority(builder, v.Priority())
		proto.TeamAddName(builder, name)
		go_a = append(go_a, proto.TeamEnd(builder))
	}
	proto.TownStartTeamsVector(builder, len(go_a))
	for _, v := range go_a {
		builder.PrependUOffsetT(v)
	}
	return builder.EndVector(len(go_a))

}

func buildDistrictTeamPermissionMap(
	builder *flatbuffers.Builder,
	target *DistrictTeamPermissionMap,
) flatbuffers.UOffsetT {

	district_ids := []uint64{}
	team_perms := []flatbuffers.UOffsetT{}
	map_count := 0
	for k, v := range target.i {
		map_count = map_count + 1
		district_ids = append(district_ids, k)
		team_perms = append(team_perms, BuildTeamPermissionMap(builder, v))
	}
	proto.DistrictTeamPermissionMapStartDistrictsVector(builder, map_count)
	for _, v := range district_ids {
		builder.PrependUint64(v)
	}
	districts := builder.EndVector(map_count)
	proto.DistrictTeamPermissionMapStartPermissionsVector(builder, map_count)
	for _, v := range team_perms {
		builder.PrependUOffsetT(v)
	}
	permissions := builder.EndVector(map_count)
	proto.DistrictTeamPermissionMapStart(builder)
	proto.DistrictTeamPermissionMapAddDistricts(builder, districts)
	proto.DistrictTeamPermissionMapAddPermissions(builder, permissions)
	return proto.DistrictTeamPermissionMapEnd(builder)
}

func buildDistrictPlayerPermissionMap(
	builder *flatbuffers.Builder,
	target *DistrictPlayerPermissionMap,
) flatbuffers.UOffsetT {
	district_ids := []uint64{}
	gamer_perms := []flatbuffers.UOffsetT{}
	map_count := 0
	for k, v := range target.i {
		map_count = map_count + 1
		district_ids = append(district_ids, k)
		gamer_perms = append(gamer_perms, BuildPlayerPermissionMap(builder, v))
	}
	proto.DistrictPlayerPermissionMapStartDistrictsVector(builder, map_count)
	for _, v := range district_ids {
		builder.PrependUint64(v)
	}
	districts := builder.EndVector(map_count)
	proto.DistrictTeamPermissionMapStartPermissionsVector(builder, map_count)
	for _, v := range gamer_perms {
		builder.PrependUOffsetT(v)
	}
	permissions := builder.EndVector(map_count)
	proto.DistrictPlayerPermissionMapStart(builder)
	proto.DistrictPlayerPermissionMapAddDistricts(builder, districts)
	proto.DistrictPlayerPermissionMapAddPermissions(builder, permissions)
	return proto.DistrictPlayerPermissionMapEnd(builder)
}

func BuildTeamPermissionMap(
	builder *flatbuffers.Builder,
	target *TeamPermissionMap,
) flatbuffers.UOffsetT {
	gp_o := BuildTeamPermissions(builder, target)
	proto.TeamPermissionMapStartPermissionsVector(builder, len(gp_o))
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
