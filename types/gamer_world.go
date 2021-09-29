package types

import (
	"errors"
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/google/uuid"

	proto "github.com/gfx-labs/etherlands/proto"
)

func (W *World) GamerCount() int {
	return len(W.gamers)
}

func (W *World) Gamers() []*Gamer {
	output := []*Gamer{}
	W.gamers_lock.RLock()
	defer W.gamers_lock.RUnlock()
	for _, v := range W.gamers {
		if v != nil {
			output = append(output, v)
		}
	}
	return output
}

func (G *Gamer) CanActIn(district *District, flag proto.AccessFlag) error {
	// first see if the district is in a town
	if district.HasTown() {
		// oh boy! first lets grab the town object
		town, err := G.W.GetTown(district.Town())
		if err != nil {
			return err
		}
		// if the gamer is a manager, they can do it.
		if town.IsManager(G) {
			return nil
		}
		// now we need to check if our member here is in any teams... this could be faster...

		// everyone starts as an outsider
		team := town.Team("outsider")
		// if the gamer is a member of our town, then we promote them to member
		if G.Town() == district.Town() {
			team = town.Team("member")
		}
		// now check if the player is in any groups
		priority := team.Priority()
		for _, v := range town.Teams() {
			if v.Has(G) {
				if v.Priority() > priority {
					team = v
				}
			}
		}
		//grab the relevant permission maps
		permission_map_district := town.DistrictTeamPermissions().
			ReadAll(district.DistrictId(), team.Name())
		permission_map_default := town.DistrictTeamPermissions().
			ReadAll(0, team.Name())
		result := proto.FlagValueNone
		if k, ok := permission_map_district[flag]; ok {
			if k == proto.FlagValueNone {
				if k, ok := permission_map_default[flag]; ok {
					result = k
				}
			} else {
				result = k
			}
		}
		if result == proto.FlagValueAllow {
			return nil
		}

		return errors.New(
			fmt.Sprintf("No permission to %s in [district.%d] of [team.%s]",
				proto.EnumNamesAccessFlag[flag],
				district.DistrictId(),
				district.Town(),
			))
	} else {
		//if not in town, grab the owner and check if they have the actioner as a friend
		// first see if the plot is linked
		uuid_str, err := G.W.Cache().GetLink(district.OwnerAddress())
		if err != nil {
			return err
		}
		uuid, err := uuid.Parse(uuid_str)
		if err != nil {
			return err
		}
		owner_obj := G.W.GetGamer(uuid)
		// yay u have a friend!!!
		if owner_obj.HasFriend(G.MinecraftId()) {
			return nil
		}
		return errors.New(
			fmt.Sprintf("No permission to %s in [district.%d]",
				proto.EnumNamesAccessFlag[flag],
				district.DistrictId()))
	}
}

func (W *World) GetGamer(gamer_id uuid.UUID) *Gamer {
	// first check if the gamer is in live cache
	W.gamers_lock.RLock()
	if val, ok := W.gamers[NewGamerKey(gamer_id)]; ok {
		W.gamers_lock.RUnlock()
		return val
	}
	//release the read lock
	W.gamers_lock.RUnlock()

	//obtain a write lock
	W.gamers_lock.Lock()
	// if not in live cache, see if gamer file exists
	if res, err := W.LoadGamer(gamer_id); err == nil {
		W.gamers_lock.Unlock()
		// add it to the cache
		W.UpdateGamer(res)
		return res
	}
	W.gamers_lock.Unlock()
	// oh no!! the gamer does not exist!!! make one!!!
	output := W.newGamer(gamer_id)
	// add it to the cache
	W.UpdateGamer(output)
	return output
}

func (W *World) newGamer(gamer_id uuid.UUID) *Gamer {
	return &Gamer{
		W:           W,
		key:         NewGamerKey(gamer_id),
		minecraftId: gamer_id,
		friends:     make(map[uuid.UUID]struct{}),
	}
}

func (W *World) LoadGamer(gamer_id uuid.UUID) (*Gamer, error) {
	bytes, err := ReadStruct("gamers", gamer_id.String())
	if err != nil {
		return nil, err
	}
	if len(bytes) < 8 {
		return nil, errors.New(fmt.Sprintf("Empty file for %s", gamer_id.String()))
	}
	read_gamer := proto.GetRootAsGamer(bytes, 0)

	read_gam := read_gamer.MinecraftId(nil)
	read_uuid := ProtoResolveUUID(read_gam)
	pending_gamer := W.newGamer(read_uuid)
	pending_gamer.town = string(read_gamer.Town())
	pending_gamer.address = string(read_gamer.Address())
	pending_gamer.nickname = string(read_gamer.Nickname())

	friend := new(proto.UUID)
	for i := 0; i < read_gamer.FriendsLength(); i++ {
		if read_gamer.Friends(friend, i) {
			pending_gamer.friends[ProtoResolveUUID(friend)] = struct{}{}
		}
	}
	return pending_gamer, nil
}

func ProtoResolveUUID(puuid *proto.UUID) uuid.UUID {
	return [16]byte{
		puuid.B0(),
		puuid.B1(),
		puuid.B2(),
		puuid.B3(),
		puuid.B4(),
		puuid.B5(),
		puuid.B6(),
		puuid.B7(),
		puuid.B8(),
		puuid.B9(),
		puuid.B10(),
		puuid.B11(),
		puuid.B12(),
		puuid.B13(),
		puuid.B14(),
		puuid.B15(),
	}
}

func (G *Gamer) Save() error {
	builder := flatbuffers.NewBuilder(1024)
	addr := builder.CreateString(G.Address())
	nick := builder.CreateString(G.Nickname())
	town := builder.CreateString(G.Town())
	gamer_friends := G.Friends()
	proto.GamerStartFriendsVector(builder, len(gamer_friends))
	for k := range gamer_friends {
		BuildUUID(builder, k)
	}
	friends_vector := builder.EndVector(len(gamer_friends))

	proto.GamerStart(builder)
	proto.GamerAddAddress(builder, addr)
	proto.GamerAddNickname(builder, nick)
	proto.GamerAddFriends(builder, friends_vector)
	uuid := proto.CreateUUID(builder, G.minecraftId[0],
		G.minecraftId[1],
		G.minecraftId[2],
		G.minecraftId[3],
		G.minecraftId[4],
		G.minecraftId[5],
		G.minecraftId[6],
		G.minecraftId[7],
		G.minecraftId[8],
		G.minecraftId[9],
		G.minecraftId[10],
		G.minecraftId[11],
		G.minecraftId[12],
		G.minecraftId[13],
		G.minecraftId[14],
		G.minecraftId[15],
	)
	proto.GamerAddMinecraftId(builder, uuid)
	proto.GamerAddTown(builder, town)

	gamer := proto.GamerEnd(builder)
	builder.Finish(gamer)

	buf := builder.FinishedBytes()
	return WriteStruct("gamers", G.MinecraftId().String(), buf)
}
