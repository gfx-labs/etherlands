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

// note - loading a gamer from memory will pause everything
// design improvement will be to make this a gamer distributor that sends to channels
// but w/e we will do that... another decade....
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
	defer W.gamers_lock.Unlock()
	// if not in live cache, see if gamer file exists
	if res, err := W.LoadGamer(gamer_id); err == nil {
		// add it to the cache
		W.UpdateGamer(res)
		return res
	}
	// oh no!! the gamer does not exist!!! make one!!!
	output := W.newGamer(gamer_id)
	// add it to the cache
	W.UpdateGamer(output)
	return output
}

func (W *World) newGamer(gamer_id uuid.UUID) *Gamer {
	return &Gamer{
		world:       W,
		key:         NewGamerKey(gamer_id),
		minecraftId: gamer_id,
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
	return &Gamer{
		world:       W,
		minecraftId: read_uuid,
		key:         NewGamerKey(read_uuid),
		town:        string(read_gamer.Town()),
		address:     string(read_gamer.Address()),
		nickname:    string(read_gamer.Nickname()),
	}, nil
}

func ProtoResolveUUID(puuid *proto.UUID) uuid.UUID {
	return [16]byte{
		byte(puuid.B0()),
		byte(puuid.B1()),
		byte(puuid.B2()),
		byte(puuid.B3()),
		byte(puuid.B4()),
		byte(puuid.B5()),
		byte(puuid.B6()),
		byte(puuid.B7()),
		byte(puuid.B8()),
		byte(puuid.B9()),
		byte(puuid.B10()),
		byte(puuid.B11()),
		byte(puuid.B12()),
		byte(puuid.B13()),
		byte(puuid.B14()),
		byte(puuid.B15()),
	}
}

func (G *Gamer) Save() error {
	builder := flatbuffers.NewBuilder(1024)
	addr := builder.CreateString(G.Address())
	nick := builder.CreateString(G.Nickname())
	town := builder.CreateString(G.Town())
	proto.GamerStart(builder)
	proto.GamerAddAddress(builder, addr)
	proto.GamerAddNickname(builder, nick)

	uuid := proto.CreateUUID(builder, int8(G.minecraftId[0]),
		int8(G.minecraftId[1]),
		int8(G.minecraftId[2]),
		int8(G.minecraftId[3]),
		int8(G.minecraftId[4]),
		int8(G.minecraftId[5]),
		int8(G.minecraftId[6]),
		int8(G.minecraftId[7]),
		int8(G.minecraftId[8]),
		int8(G.minecraftId[9]),
		int8(G.minecraftId[10]),
		int8(G.minecraftId[11]),
		int8(G.minecraftId[12]),
		int8(G.minecraftId[13]),
		int8(G.minecraftId[14]),
		int8(G.minecraftId[15]),
	)
	proto.GamerAddMinecraftId(builder, uuid)
	proto.GamerAddTown(builder, town)

	gamer := proto.GamerEnd(builder)
	builder.Finish(gamer)

	buf := builder.FinishedBytes()
	return WriteStruct("gamers", G.MinecraftId().String(), buf)
}
