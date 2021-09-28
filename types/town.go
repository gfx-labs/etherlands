package types

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Team struct {
	name     string
	town     string
	priority int64
	members  map[uuid.UUID]struct{}

	sync.RWMutex
}

func (G *Team) Name() string {
	return G.name
}

func (G *Team) Town() string {
	return G.town
}

func (G *Team) SetPriority(newPriority int64) {
	if G.name == "manager" || G.name == "outsider" || G.name == "member" {
		return
	}
	G.Lock()
	if newPriority < 0 {
		G.priority = 0
	} else if newPriority > 100 {
		G.priority = 100
	} else {
		G.priority = newPriority
	}
	G.Unlock()
}

func (G *Team) Priority() int64 {
	if G.name == "manager" {
		return 100
	}
	if G.name == "outsider" {
		return -100
	}
	if G.name == "member" {
		return -1
	}
	G.RLock()
	defer G.RUnlock()
	return G.priority
}

func (G *Team) Has(gamer *Gamer) bool {
	G.RLock()
	defer G.RUnlock()
	_, ok := G.members[gamer.MinecraftId()]
	return ok
}

func (G *Team) add(gamer *Gamer) {
	G.RLock()
	defer G.RUnlock()
	G.members[gamer.MinecraftId()] = struct{}{}
}

func (G *Team) remove(gamer *Gamer) {
	G.RLock()
	defer G.RUnlock()
	delete(G.members, gamer.MinecraftId())
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

func (T *Town) TeamAddMember(manager *Gamer, team_name string, gamer *Gamer) error {
	if team_name == "member" || team_name == "outsider" {
		return errors.New("Cannot add player to default group")
	}
	if team_name == "manager" && manager.MinecraftId() != T.Owner() {
		return errors.New("Only the owner may add to the manager group")
	}
	if !T.CanAction(manager, gamer) {
		return errors.New("You may not adjust another managers groups")
	}
	team := T.Team(team_name)
	if team == nil {
		return errors.New("That team does not exist")
	}
	team.add(gamer)
	return nil
}

func (T *Town) TeamRemoveMember(manager *Gamer, team_name string, gamer *Gamer) error {
	if team_name == "member" || team_name == "outsider" {
		return errors.New("Cannot add player to default group")
	}
	if team_name == "manager" && manager.MinecraftId() != T.Owner() {
		return errors.New("only the owner may remove from the manager group")
	}
	if !T.CanAction(manager, gamer) {
		return errors.New("You may not adjust another managers groups")
	}
	team := T.Team(team_name)
	if team == nil {
		return errors.New("That team does not exist")
	}
	team.remove(gamer)
	return nil
}

func (T *Town) addTeam(name string) error {
	if name == "member" || name == "outsider" || name == "manager" {
		return errors.New("Cannot create a team using a default name")
	}
	T.Lock()
	defer T.Unlock()
	if _, ok := T.teams[name]; !ok {
		T.teams[name] = &Team{
			name: name,
			town: T.name,
		}
		return nil
	}
	return errors.New(fmt.Sprintf("Team with name %s already exists", name))
}

func (T *Town) removeTeam(name string) error {
	if name == "member" || name == "outsider" || name == "manager" {
		return errors.New("Cannot remove a team with default name")
	}
	T.Lock()
	defer T.Unlock()
	if _, ok := T.teams[name]; ok {
		delete(T.teams, name)
		return nil
	}
	return errors.New(fmt.Sprintf("Team with name %s doesn't exist", name))
}

type Town struct {
	name string

	owner uuid.UUID

	teams     map[string]*Team
	districts []uint64

	defaultPermissions *TeamPermissionMap

	districtPlayerPermissions *DistrictPlayerPermissionMap
	districtTeamPermissions   *DistrictTeamPermissionMap
	district_player_lock      *MapLock
	district_team_lock        *MapLock

	sync.RWMutex

	W *World

	invites     map[uuid.UUID]time.Time
	invite_lock sync.Mutex
}

func (T *Town) CreateTeam(manager *Gamer, name string) error {
	if T.IsManager(manager) {
		return T.addTeam(name)
	}
	return errors.New("You must be a manager to create a team")
}

func (T *Town) RemoveTeam(manager *Gamer, name string) error {
	if T.IsManager(manager) {
		return T.removeTeam(name)
	}
	return errors.New("You must be a manager to create a team")
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

func (T *Town) DefaultPermissions() *TeamPermissionMap {
	T.RLock()
	defer T.RUnlock()
	return T.defaultPermissions
}

func (T *Town) DistrictPlayerPermissions() *DistrictPlayerPermissionMap {
	T.RLock()
	defer T.RUnlock()
	return T.districtPlayerPermissions
}

func (T *Town) DistrictTeamPermissions() *DistrictTeamPermissionMap {
	T.RLock()
	defer T.RUnlock()
	return T.districtTeamPermissions
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
