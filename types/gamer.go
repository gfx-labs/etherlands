package types

import (
	"strings"
	"sync"

	proto "github.com/gfx-labs/etherlands/proto"
	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/google/uuid"
)

type Gamer struct {
	nickname    string
	address     string
	town        string
	minecraftId uuid.UUID
	mutex       sync.RWMutex

	key   FamilyKey
	world *World

	pos_x     int64
	pos_y     int64
	pos_z     int64
	pos_mutex sync.RWMutex
}

func NewGamerKey(gamer_id uuid.UUID) FamilyKey {
	return FamilyKey{
		datatype: GAMER_FAMILY,
		subkey:   strings.ToLower(gamer_id.String()),
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

func (G *Gamer) GetTown() string {
	G.mutex.RLock()
	defer G.mutex.RUnlock()
	return G.town
}

func (G *Gamer) SetTown(name string) {
	G.mutex.RLock()
	defer G.mutex.RUnlock()
	G.town = name
}

func (G *Gamer) HasTown() bool {
	G.mutex.RLock()
	defer G.mutex.RUnlock()
	return G.world.TownOfGamer(G) != ""
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

func (G *Gamer) SetPosXYZ(x, y, z int64) {
	G.mutex.Lock()
	defer G.mutex.Unlock()
	G.pos_x = x
	G.pos_y = y
	G.pos_z = z
}
func (G *Gamer) GetPosXYZ() (x, y, z int64) {
	G.mutex.Lock()
	defer G.mutex.Unlock()
	return G.pos_x, G.pos_y, G.pos_z
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
