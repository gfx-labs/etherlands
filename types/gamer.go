package types

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Gamer struct {
	nickname string
	address  string
	town     string

	friends map[uuid.UUID]struct{}

	minecraftId uuid.UUID
	mutex       sync.RWMutex

	key FamilyKey
	W   *World

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

func (G *Gamer) JoinTown(town *Town) error {
	if town.CheckInvite(G, time.Minute*15) {
		G.SetTown(town.Name())
		G.Update()
		return nil
	}
	return errors.New(fmt.Sprintf("You must be invited to join [town.%s]", town.Name()))
}

func (G *Gamer) DeleteTown(town *Town, validate string) error {
	if town.Name() != validate {
		return errors.New(
			fmt.Sprintf(
				"Run command with argument %s (e.g. /town delete %s)",
				town.Name(),
				town.Name(),
			),
		)
	}
	if town.Owner() == G.MinecraftId() {
		if len(G.W.GamersOfTown(town.Name())) == 1 {
			G.SetTown("")
			G.Update()
			G.W.DeleteTown(town)
			return nil
		}
		return errors.New(fmt.Sprintf("kick everyone else from your town first"))
	}
	return errors.New(fmt.Sprintf("only owner may delete team"))
}

func (G *Gamer) LeaveTown(town *Town) error {
	if G.Town() != "" {
		if town.Owner() != G.minecraftId {
			G.SetTown("")
			G.Update()
			return nil
		}
		return errors.New(fmt.Sprintf("You may not leave the team you own"))
	}
	return errors.New(fmt.Sprintf("You already are in no town"))
}

func (G *Gamer) KickTown(kicked *Gamer, town *Town) error {
	if kicked.Town() == G.Town() {
		if town.CanAction(G, kicked) {
			kicked.SetTown("")
			kicked.Update()
			return nil
		}
		return errors.New(
			fmt.Sprintf(
				"You do not have permission to kick [uuid.%s]",
				kicked.MinecraftId().String(),
			),
		)
	}
	return errors.New(
		fmt.Sprintf("[uuid.%s] is not a member of your town", kicked.MinecraftId().String()),
	)
}

func (G *Gamer) Friends() map[uuid.UUID]struct{} {
	G.mutex.RLock()
	defer G.mutex.RUnlock()
	return G.friends
}
func (G *Gamer) AddFriend(target uuid.UUID) {
	G.mutex.RLock()
	defer G.mutex.RUnlock()
	G.friends[target] = struct{}{}
}
func (G *Gamer) HasFriend(target uuid.UUID) bool {
	G.mutex.RLock()
	defer G.mutex.RUnlock()
	_, ok := G.friends[target]
	return ok
}

func (G *Gamer) RemoveFriend(target uuid.UUID) {
	G.mutex.RLock()
	defer G.mutex.RUnlock()
	delete(G.friends, target)
}

func (G *Gamer) Nickname() string {
	G.mutex.RLock()
	defer G.mutex.RUnlock()
	return G.minecraftId.String()
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
	G.address = strings.ToLower(address)
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
