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
type TeamPermissionMap struct {
	i       map[string]map[proto.AccessFlag]proto.FlagValue
	mutexes sync.Map
}

func NewTeamPermissionMap() *TeamPermissionMap {
	return &TeamPermissionMap{
		i: make(map[string]map[proto.AccessFlag]proto.FlagValue),
	}
}

func NewPlayerPermissionMap() *PlayerPermissionMap {
	return &PlayerPermissionMap{
		i: make(map[uuid.UUID]map[proto.AccessFlag]proto.FlagValue),
	}
}

type MapLock struct {
	sync.Map
}

func (D *MapLock) lock(district_id uint64) func() {
	value, _ := D.LoadOrStore(district_id, &sync.Mutex{})
	mtx := value.(*sync.Mutex)
	mtx.Lock()
	return func() { mtx.Unlock() }
}

func NewMapLock() *MapLock {
	return &MapLock{}
}

func (P *TeamPermissionMap) lock(name string) func() {
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

func (P *TeamPermissionMap) insert(
	team_id string,
	flag proto.AccessFlag,
	value proto.FlagValue,
) {
	unlock := P.lock(team_id)
	defer unlock()
	if _, ok := P.i[team_id]; !ok {
		P.i[team_id] = make(map[proto.AccessFlag]proto.FlagValue)
	}
	P.i[team_id][flag] = value
}

func (P *TeamPermissionMap) read(
	team_id string,
	flag proto.AccessFlag,
) proto.FlagValue {
	unlock := P.lock(team_id)
	defer unlock()
	if v, ok := P.i[team_id]; ok {
		if v2, ok := v[flag]; ok {
			return v2
		}
	}
	return proto.FlagValueNone
}
