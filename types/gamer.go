package types

import (
	"strings"
	"sync"

	"github.com/google/uuid"
)

type Gamer struct {
	nickname string
	address  string
	town     string

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

func (G *Gamer) Town() string {
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
	return G.town != ""
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
