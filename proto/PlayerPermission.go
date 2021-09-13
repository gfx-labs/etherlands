// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Etherlands

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type PlayerPermission struct {
	_tab flatbuffers.Table
}

func GetRootAsPlayerPermission(buf []byte, offset flatbuffers.UOffsetT) *PlayerPermission {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &PlayerPermission{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *PlayerPermission) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PlayerPermission) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *PlayerPermission) MinecraftId(obj *UUID) *UUID {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(UUID)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *PlayerPermission) Flag() AccessFlag {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return AccessFlag(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *PlayerPermission) MutateFlag(n AccessFlag) bool {
	return rcv._tab.MutateByteSlot(6, byte(n))
}

func (rcv *PlayerPermission) Value() FlagValue {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return FlagValue(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *PlayerPermission) MutateValue(n FlagValue) bool {
	return rcv._tab.MutateByteSlot(8, byte(n))
}

func PlayerPermissionStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func PlayerPermissionAddMinecraftId(builder *flatbuffers.Builder, minecraftId flatbuffers.UOffsetT) {
	builder.PrependStructSlot(0, flatbuffers.UOffsetT(minecraftId), 0)
}
func PlayerPermissionAddFlag(builder *flatbuffers.Builder, flag AccessFlag) {
	builder.PrependByteSlot(1, byte(flag), 0)
}
func PlayerPermissionAddValue(builder *flatbuffers.Builder, value FlagValue) {
	builder.PrependByteSlot(2, byte(value), 0)
}
func PlayerPermissionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
