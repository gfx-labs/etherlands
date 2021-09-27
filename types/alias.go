package types

import (
	"sync"

	proto "github.com/gfx-labs/etherlands/proto"
	uuid "github.com/google/uuid"
)

type PlayerPermissionMap struct {
	i       map[uuid.UUID]map[proto.AccessFlag]proto.FlagValue
	mutexes sync.Map
}
type GroupPermissionMap struct {
	i       map[string]map[proto.AccessFlag]proto.FlagValue
	mutexes sync.Map
}

type DistrictLock struct {
	sync.Map
}

func (D *DistrictLock) lock(district_id uint64) func() {
	value, _ := D.LoadOrStore(district_id, &sync.Mutex{})
	mtx := value.(*sync.Mutex)
	mtx.Lock()
	return func() { mtx.Unlock() }
}

func NewDistrictLock() *DistrictLock {
	return &DistrictLock{}
}

func NewGroupPermissionMap() *GroupPermissionMap {
	return &GroupPermissionMap{
		i: make(map[string]map[proto.AccessFlag]proto.FlagValue),
	}
}

func NewPlayerPermissionMap() *PlayerPermissionMap {
	return &PlayerPermissionMap{
		i: make(map[uuid.UUID]map[proto.AccessFlag]proto.FlagValue),
	}
}

func (P *GroupPermissionMap) lock(name string) func() {
	value, _ := P.mutexes.LoadOrStore(name, &sync.Mutex{})
	mtx := value.(*sync.Mutex)
	mtx.Lock()
	return func() { mtx.Unlock() }
}
func (P *PlayerPermissionMap) lock(name uuid.UUID) func() {
	value, _ := P.mutexes.LoadOrStore(name, &sync.Mutex{})
	mtx := value.(*sync.Mutex)
	mtx.Lock()
	return func() { mtx.Unlock() }
}

func (P *PlayerPermissionMap) insert(
	gamer_id uuid.UUID,
	flag proto.AccessFlag,
	value proto.FlagValue,
) {
	unlock := P.lock(gamer_id)
	defer unlock()
	if _, ok := P.i[gamer_id]; !ok {
		P.i[gamer_id] = make(map[proto.AccessFlag]proto.FlagValue)
	}
	P.i[gamer_id][flag] = value
}

func (P *PlayerPermissionMap) read(
	gamer_id uuid.UUID,
	flag proto.AccessFlag,
) proto.FlagValue {
	unlock := P.lock(gamer_id)
	defer unlock()
	if v, ok := P.i[gamer_id]; ok {
		if v2, ok := v[flag]; ok {
			return v2
		}
	}
	return proto.FlagValueNone
}

func (P *GroupPermissionMap) insert(
	group_id string,
	flag proto.AccessFlag,
	value proto.FlagValue,
) {
	unlock := P.lock(group_id)
	defer unlock()
	if _, ok := P.i[group_id]; !ok {
		P.i[group_id] = make(map[proto.AccessFlag]proto.FlagValue)
	}
	P.i[group_id][flag] = value
}

func (P *GroupPermissionMap) read(
	group_id string,
	flag proto.AccessFlag,
) proto.FlagValue {
	unlock := P.lock(group_id)
	defer unlock()
	if v, ok := P.i[group_id]; ok {
		if v2, ok := v[flag]; ok {
			return v2
		}
	}
	return proto.FlagValueNone
}
