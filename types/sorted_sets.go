package types

import "github.com/google/uuid"

func (W *World) PlotsOfDistrict(district_id uint64) []uint64 {
	W.cache.plot_lock.RLock()
	defer W.cache.plot_lock.RUnlock()
	return W.cache.plot_district.GetKeysByScore(district_id)
}

func (W *World) DistrictOfOwner(address string) []uint64 {
	W.cache.district_lock.RLock()
	defer W.cache.district_lock.RUnlock()
	return W.cache.district_owner.GetKeysByScore(address)
}

func (W *World) TownOfGamer(gamer *Gamer) string {
	W.cache.uuid_town_lock.Lock()
	defer W.cache.uuid_town_lock.Unlock()
	item := W.cache.uuid_town.GetByKey(gamer.MinecraftId())
	if item == nil {
		return ""
	}
	return item.Score()

}
func (W *World) GamersOfTown(name string) map[uuid.UUID]struct{} {
	W.cache.uuid_town_lock.Lock()
	defer W.cache.uuid_town_lock.Unlock()
	return W.cache.uuid_town.GetKeysByScore(name)
}
