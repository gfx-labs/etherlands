package types

import (
	"sync"

	"github.com/google/uuid"
)

type Group struct {
	name    string
	town    string
	members []uuid.UUID

	sync.RWMutex
}

func (G *Group) Name() string {
	G.RLock()
	defer G.RUnlock()
	return G.name
}

func (G *Group) Town() string {
	G.RLock()
	defer G.RUnlock()
	return G.town
}

func (G *Group) Members() []uuid.UUID {
	G.RLock()
	defer G.RUnlock()
	return G.members
}

type Town struct {
	name string

	owner   uuid.UUID
	members map[uuid.UUID]struct{}

	groups    map[string]*Group
	districts []uint64

	defaultPlayerPermissions *PlayerPermissionMap
	defaultGroupPermissions  *GroupPermissionMap

	districtPlayerPermissions map[uint64]*PlayerPermissionMap
	districtGroupPermissions  map[uint64]*GroupPermissionMap
	district_player_lock      *DistrictLock
	district_group_lock       *DistrictLock

	sync.RWMutex

	W *World
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

func (T *Town) DefaultGroupPermissions() *GroupPermissionMap {
	T.RLock()
	defer T.RUnlock()
	return T.defaultGroupPermissions
}

func (T *Town) DistrictPlayerPermissions() map[uint64]*PlayerPermissionMap {
	T.RLock()
	defer T.RUnlock()
	return T.districtPlayerPermissions
}

func (T *Town) DistrictGroupPermissions() map[uint64]*GroupPermissionMap {
	T.RLock()
	defer T.RUnlock()
	return T.districtGroupPermissions
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

func (T *Town) DistrictGroupPermission(district_id uint64) *GroupPermissionMap {
	T.RLock()
	defer T.RUnlock()
	if v, ok := T.districtGroupPermissions[district_id]; ok {
		return v
	}
	T.districtGroupPermissions[district_id] = NewGroupPermissionMap()
	return T.districtGroupPermissions[district_id]
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

func (T *Town) Members() map[uuid.UUID]struct{} {
	T.RLock()
	defer T.RUnlock()
	return T.members
}

func (T *Town) AddMember(id uuid.UUID) {
	T.Lock()
	defer T.Unlock()
	T.members[id] = struct{}{}
}

func (T *Town) Groups() []*Group {
	T.RLock()
	defer T.RUnlock()
	return T.Groups()
}

func (T *Town) Districts() []uint64 {
	T.RLock()
	defer T.RUnlock()
	return T.districts
}
