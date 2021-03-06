// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Etherlands

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Plot struct {
	_tab flatbuffers.Table
}

func GetRootAsPlot(buf []byte, offset flatbuffers.UOffsetT) *Plot {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Plot{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Plot) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Plot) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Plot) PlotId() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Plot) MutatePlotId(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *Plot) DistrictId() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Plot) MutateDistrictId(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func (rcv *Plot) X() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Plot) MutateX(n int64) bool {
	return rcv._tab.MutateInt64Slot(8, n)
}

func (rcv *Plot) Z() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Plot) MutateZ(n int64) bool {
	return rcv._tab.MutateInt64Slot(10, n)
}

func PlotStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func PlotAddPlotId(builder *flatbuffers.Builder, plotId uint64) {
	builder.PrependUint64Slot(0, plotId, 0)
}
func PlotAddDistrictId(builder *flatbuffers.Builder, districtId uint64) {
	builder.PrependUint64Slot(1, districtId, 0)
}
func PlotAddX(builder *flatbuffers.Builder, x int64) {
	builder.PrependInt64Slot(2, x, 0)
}
func PlotAddZ(builder *flatbuffers.Builder, z int64) {
	builder.PrependInt64Slot(3, z, 0)
}
func PlotEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
