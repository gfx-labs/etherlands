package types

func (W *World) PlotsOfDistrict(district_id uint64) []uint64 {
	return W.plot_district.GetKeysByScore(district_id)
}

func (W *World) DistrictOfOwner(address string) []uint64 {
	return W.district_owner.GetKeysByScore(address)
}
