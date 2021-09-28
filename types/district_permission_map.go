package types

import (
	"sync"

	proto "github.com/gfx-labs/etherlands/proto"
	"github.com/google/uuid"
)

type DistrictPlayerPermissionMap struct {
	i       map[uint64]*PlayerPermissionMap
	mutexes sync.Map
	global  sync.RWMutex
}
type DistrictTeamPermissionMap struct {
	i       map[uint64]*TeamPermissionMap
	mutexes sync.Map
	global  sync.RWMutex
}

func (P *DistrictTeamPermissionMap) lock(id uint64) func() {
	P.global.RLock()
	defer P.global.RUnlock()
	value, _ := P.mutexes.LoadOrStore(id, &sync.Mutex{})
	mtx := value.(*sync.Mutex)
	mtx.Lock()
	return func() { mtx.Unlock() }
}
func (P *DistrictPlayerPermissionMap) lock(id uint64) func() {
	P.global.RLock()
	defer P.global.RUnlock()
	value, _ := P.mutexes.LoadOrStore(id, &sync.Mutex{})
	mtx := value.(*sync.Mutex)
	mtx.Lock()
	return func() { mtx.Unlock() }
}

func NewDistrictTeamPermissionMap() *DistrictTeamPermissionMap {
	return &DistrictTeamPermissionMap{
		i: make(map[uint64]*TeamPermissionMap),
	}
}

func NewDistrictPlayerPermissionMap() *DistrictPlayerPermissionMap {
	return &DistrictPlayerPermissionMap{
		i: make(map[uint64]*PlayerPermissionMap),
	}
}

func (P *DistrictPlayerPermissionMap) Insert(
	district_id uint64,
	gamer_id uuid.UUID,
	flag proto.AccessFlag,
	value proto.FlagValue,
) {
	unlock := P.lock(district_id)
	defer unlock()
	if _, ok := P.i[district_id]; !ok {
		P.i[district_id] = NewPlayerPermissionMap()
	}
	P.i[district_id].insert(gamer_id, flag, value)
}

func (P *DistrictPlayerPermissionMap) Read(
	district_id uint64,
	gamer_id uuid.UUID,
	flag proto.AccessFlag,
) proto.FlagValue {
	unlock := P.lock(district_id)
	defer unlock()
	if v, ok := P.i[district_id]; ok {
		return v.read(gamer_id, flag)
	}
	return proto.FlagValueNone
}

func (P *DistrictTeamPermissionMap) Insert(
	district_id uint64,
	team_id string,
	flag proto.AccessFlag,
	value proto.FlagValue,
) {
	unlock := P.lock(district_id)
	defer unlock()
	if _, ok := P.i[district_id]; !ok {
		P.i[district_id] = NewTeamPermissionMap()
	}
	P.i[district_id].insert(team_id, flag, value)
}

func (P *DistrictTeamPermissionMap) Read(
	district_id uint64,
	team_id string,
	flag proto.AccessFlag,
) proto.FlagValue {
	unlock := P.lock(district_id)
	defer unlock()
	if v, ok := P.i[district_id]; ok {
		return v.read(team_id, flag)
	}
	return proto.FlagValueNone
}
