package types

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Team struct {
	name    string
	town    string
	members map[uuid.UUID]struct{}

	sync.RWMutex
}

func (G *Team) Name() string {
	G.RLock()
	defer G.RUnlock()
	return G.name
}

func (G *Team) Town() string {
	G.RLock()
	defer G.RUnlock()
	return G.town
}

func (G *Team) Has(gamer *Gamer) bool {
	G.RLock()
	defer G.RUnlock()
	_, ok := G.members[gamer.MinecraftId()]
	return ok
}

func (G *Team) Members() map[uuid.UUID]struct{} {
	G.RLock()
	defer G.RUnlock()
	return G.members
}
func (T *Town) Team(name string) *Team {
	T.RLock()
	defer T.RUnlock()
	if v, ok := T.teams[name]; ok {
		return v
	}
	return nil
}

type Town struct {
	name string

	owner uuid.UUID

	teams     map[string]*Team
	districts []uint64

	defaultPlayerPermissions *PlayerPermissionMap
	defaultTeamPermissions   *TeamPermissionMap

	districtPlayerPermissions map[uint64]*PlayerPermissionMap
	districtTeamPermissions   map[uint64]*TeamPermissionMap
	district_player_lock      *DistrictLock
	district_team_lock        *DistrictLock

	sync.RWMutex

	W *World

	invites     map[uuid.UUID]time.Time
	invite_lock sync.Mutex
}

func (T *Town) CheckInvite(gamer *Gamer, timeout time.Duration) bool {
	T.invite_lock.Lock()
	defer T.invite_lock.Unlock()
	if v, ok := T.invites[gamer.MinecraftId()]; ok {
		if (time.Now().Sub(v)) > timeout {
			delete(T.invites, gamer.MinecraftId())
			return false
		} else {
			delete(T.invites, gamer.MinecraftId())
			return true
		}
	}
	return false
}

func (T *Town) ProcessInvites(interval time.Duration) {
	for {
		time.Sleep(interval)
		T.invite_lock.Lock()
		for k, v := range T.invites {
			if (time.Now().Sub(v)) > interval {
				delete(T.invites, k)
			}
		}
		T.invite_lock.Unlock()
	}
}

func (T *Town) InviteGamer(sender, receiver *Gamer) error {
	if T.IsManager(sender) {
		T.invite_lock.Lock()
		T.invites[receiver.MinecraftId()] = time.Now()
		T.invite_lock.Unlock()
		return nil
	}
	return errors.New("you must be a manager to invite")
}

func (T *Town) IsManager(gamer *Gamer) bool {
	if team := T.Team("manager"); team != nil {
		if team.Has(gamer) {
			return true
		}
	}
	return T.Owner() == gamer.MinecraftId()
}

func (T *Town) CanAction(actor *Gamer, target *Gamer) bool {
	if actor.MinecraftId() == T.Owner() {
		return true
	}
	if team := T.Team("manager"); team != nil {
		if team.Has(actor) && !team.Has(target) {
			return true
		}
	}
	return false
}

func (T *Town) GetKey() FamilyKey {
	T.RLock()
	defer T.RUnlock()
	return NewTownKey(T.name)
}

func NewTownKey(town_name string) FamilyKey {
	return FamilyKey{
		datatype: TOWN_FAMILY,
		subkey:   town_name,
	}
}

func (T *Town) DefaultPlayerPermissions() *PlayerPermissionMap {
	T.RLock()
	defer T.RUnlock()
	return T.defaultPlayerPermissions
}

func (T *Town) DefaultTeamPermissions() *TeamPermissionMap {
	T.RLock()
	defer T.RUnlock()
	return T.defaultTeamPermissions
}

func (T *Town) DistrictPlayerPermissions() map[uint64]*PlayerPermissionMap {
	T.RLock()
	defer T.RUnlock()
	return T.districtPlayerPermissions
}

func (T *Town) DistrictTeamPermissions() map[uint64]*TeamPermissionMap {
	T.RLock()
	defer T.RUnlock()
	return T.districtTeamPermissions
}

func (T *Town) DistrictPlayerPermission(district_id uint64) *PlayerPermissionMap {
	T.RLock()
	defer T.RUnlock()
	if v, ok := T.districtPlayerPermissions[district_id]; ok {
		return v
	}
	T.districtPlayerPermissions[district_id] = NewPlayerPermissionMap()
	return T.districtPlayerPermissions[district_id]
}

func (T *Town) DistrictTeamPermission(district_id uint64) *TeamPermissionMap {
	T.RLock()
	defer T.RUnlock()
	if v, ok := T.districtTeamPermissions[district_id]; ok {
		return v
	}
	T.districtTeamPermissions[district_id] = NewTeamPermissionMap()
	return T.districtTeamPermissions[district_id]
}

func (T *Town) Owner() uuid.UUID {
	T.RLock()
	defer T.RUnlock()
	return T.owner
}

func (T *Town) SetOwner(owner uuid.UUID) {
	T.Lock()
	defer T.Unlock()
	T.owner = owner
}

func (T *Town) Name() string {
	T.RLock()
	defer T.RUnlock()
	return T.name
}

func (T *Town) RemoveName() {
	T.Lock()
	defer T.Unlock()
	T.name = ""
}
func (T *Town) Members() map[uuid.UUID]struct{} {
	return T.W.GamersOfTown(T.Name())
}

func (T *Town) Teams() map[string]*Team {
	T.RLock()
	defer T.RUnlock()
	return T.teams
}

func (T *Town) Districts() []uint64 {
	T.RLock()
	defer T.RUnlock()
	return T.districts
}
