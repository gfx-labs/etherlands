package types

import (
	"sync"

	proto "github.com/gfx-labs/etherlands/proto"
	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/google/uuid"
)

type Gamer struct {
	nickname    string
	address     string
	minecraftId uuid.UUID

	mutex sync.RWMutex

	key   FamilyKey
	world *World
}

func NewGamerKey(gamer_id uuid.UUID) FamilyKey {
	return FamilyKey{
		datatype: GAMER_FAMILY,
		subkey:   gamer_id.String(),
	}
}

func (G *Gamer) Nickname() string {
	G.mutex.RLock()
	defer G.mutex.RUnlock()
	return G.nickname
}

func (G *Gamer) Address() string {
	G.mutex.RLock()
	defer G.mutex.RUnlock()
	return G.address
}

func (G *Gamer) MinecraftId() uuid.UUID {
	G.mutex.RLock()
	defer G.mutex.RUnlock()
	return G.minecraftId
}

func (G *Gamer) SetAddress(address string) {
	G.mutex.Lock()
	defer G.mutex.Unlock()
	G.address = address
}

func (G *Gamer) GetKey() FamilyKey {
	return G.key
}

func (G *Gamer) Save() error {
	builder := flatbuffers.NewBuilder(1024)
	addr := builder.CreateString(G.Address())
	nick := builder.CreateString(G.Nickname())
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

	gamer := proto.GamerEnd(builder)
	builder.Finish(gamer)

	buf := builder.FinishedBytes()
	return WriteStruct("gamers", G.MinecraftId().String(), buf)
}
