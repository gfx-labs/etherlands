// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Etherlands

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TeamPermission struct {
	_tab flatbuffers.Table
}

func GetRootAsTeamPermission(buf []byte, offset flatbuffers.UOffsetT) *TeamPermission {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TeamPermission{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TeamPermission) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TeamPermission) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TeamPermission) Team() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *TeamPermission) Flag() AccessFlag {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return AccessFlag(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *TeamPermission) MutateFlag(n AccessFlag) bool {
	return rcv._tab.MutateByteSlot(6, byte(n))
}

func (rcv *TeamPermission) Value() FlagValue {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return FlagValue(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *TeamPermission) MutateValue(n FlagValue) bool {
	return rcv._tab.MutateByteSlot(8, byte(n))
}

func TeamPermissionStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func TeamPermissionAddTeam(builder *flatbuffers.Builder, team flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(team), 0)
}
func TeamPermissionAddFlag(builder *flatbuffers.Builder, flag AccessFlag) {
	builder.PrependByteSlot(1, byte(flag), 0)
}
func TeamPermissionAddValue(builder *flatbuffers.Builder, value FlagValue) {
	builder.PrependByteSlot(2, byte(value), 0)
}
func TeamPermissionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
