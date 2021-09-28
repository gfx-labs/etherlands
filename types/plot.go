package types

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	proto "github.com/gfx-labs/etherlands/proto"
	flatbuffers "github.com/google/flatbuffers/go"
)

type Plot struct {
	W           *World
	x           int64
	z           int64
	plot_id     uint64
	district_id uint64
	mutex       sync.RWMutex

	key FamilyKey
}

func (P *Plot) X() int64 {
	return P.x
}

func (P *Plot) Z() int64 {
	return P.z
}

func (P *Plot) PlotId() uint64 {
	return P.plot_id
}

func (P *Plot) DistrictId() uint64 {
	P.mutex.RLock()
	defer P.mutex.RUnlock()
	return P.district_id
}

func (P *Plot) SetDistrictId(id uint64) {
	P.mutex.Lock()
	defer P.mutex.Unlock()
	P.district_id = id
}

func (P *Plot) GetLocation() [2]int64 {
	return [2]int64{P.x, P.z}
}

func (P *Plot) GetKey() FamilyKey {
	return P.key
}

func (W *World) newPlot(x, z int64, plotId, districtId uint64) *Plot {
	output := &Plot{
		W:           W,
		plot_id:     plotId,
		district_id: districtId,
		key:         NewPlotKey(plotId),
		x:           x,
		z:           z,
	}
	return output
}

func (W *World) LoadPlot(chain_id uint64) (*Plot, error) {
	bytes, err := ReadStruct("plots", strconv.FormatUint(chain_id, 10))
	if err != nil {
		return nil, err
	}
	if len(bytes) < 8 {
		return nil, errors.New(fmt.Sprintf("Empty file for %d", chain_id))
	}
	read_plot := proto.GetRootAsPlot(bytes, 0)
	return W.newPlot(
		read_plot.X(),
		read_plot.Z(),
		read_plot.PlotId(),
		read_plot.DistrictId(),
	), nil
}

func (P *Plot) Save() error {
	builder := flatbuffers.NewBuilder(1024)
	proto.PlotStart(builder)
	proto.PlotAddPlotId(builder, P.PlotId())
	proto.PlotAddDistrictId(builder, P.DistrictId())
	proto.PlotAddX(builder, P.X())
	proto.PlotAddZ(builder, P.Z())
	plot := proto.PlotEnd(builder)
	builder.Finish(plot)

	buf := builder.FinishedBytes()
	return WriteStruct("plots", strconv.FormatUint(P.PlotId(), 10), buf)
}

func NewPlotKey(id uint64) FamilyKey {
	return FamilyKey{
		datatype: PLOT_FAMILY,
		subkey:   strconv.FormatUint(id, 10),
	}
}
