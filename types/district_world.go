package types

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	proto "github.com/gfx-labs/etherlands/proto"
)

func (W *World) DistrictCount() int {
	return len(W.districts)
}

func (W *World) Districts() []*District {
	output := []*District{}
	W.districts_lock.RLock()
	defer W.districts_lock.RUnlock()
	for _, v := range W.districts {
		if v != nil {
			output = append(output, v)
		}
	}
	return output
}

func (D *District) Plots() map[uint64]*Plot {
	D.mutex.RLock()
	defer D.mutex.RUnlock()
	output := make(map[uint64]*Plot)
	for _, v := range D.W.PlotsOfDistrict(D.DistrictId()) {
		plot, err := D.W.GetPlot(v)
		if err == nil {
			output[v] = plot
		}
	}
	return output
}

func (W *World) GetDistrict(district_id uint64) (*District, error) {
	W.districts_lock.RLock()
	defer W.districts_lock.RUnlock()
	if val, ok := W.districts[NewDistrictKey(district_id)]; ok {
		return val, nil
	}
	return nil, errors.New(fmt.Sprintf("district %d could not be found", district_id))
}

func (W *World) newDistrict(id uint64, ownerAddress string, nickname [24]byte) *District {
	output := &District{
		W:             W,
		district_id:   id,
		owner_address: strings.ToLower(ownerAddress),
		nickname:      &nickname,
		key:           NewDistrictKey(id),
	}
	W.UpdateDistrict(output)
	return output
}

func (W *World) LoadDistrict(chain_id uint64) (*District, error) {
	bytes, err := ReadStruct("districts", strconv.FormatUint(chain_id, 10))
	if err != nil {
		return nil, err
	}
	if len(bytes) < 8 {
		return nil, errors.New(fmt.Sprintf("Empty file for %d", chain_id))
	}
	read_district := proto.GetRootAsDistrict(bytes, 0)
	fixed_name := [24]byte{}
	for i := 0; (i < read_district.NicknameLength()) && (i < 24); i++ {
		if i >= read_district.NicknameLength() {
			fixed_name[i] = 0
		} else {
			fixed_name[i] = byte(read_district.Nickname(i))
		}
	}

	town_name := string(read_district.Town())

	out := W.newDistrict(
		read_district.ChainId(),
		string(read_district.OwnerAddress()),
		fixed_name,
	)
	out.setTown(town_name)
	return out
}
