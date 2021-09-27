package types

import "github.com/google/uuid"

func (W *World) PlotsOfDistrict(district_id uint64) []uint64 {
	return W.plot_district.GetKeysByScore(district_id)
}

func (W *World) DistrictOfOwner(address string) []uint64 {
	return W.district_owner.GetKeysByScore(address)
}

func (W *World) TownOfGamer(gamer *Gamer) string {
	return W.uuid_town.GetByKey(gamer.MinecraftId()).Score()
}

func (W *World) GamersOfTown(name string) []uuid.UUID {
	return W.uuid_town.GetKeysByScore(name)
}
