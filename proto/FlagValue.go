// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package Etherlands

import "strconv"

type FlagValue byte

const (
	FlagValueNone  FlagValue = 0
	FlagValueAllow FlagValue = 1
	FlagValueDeny  FlagValue = 2
)

var EnumNamesFlagValue = map[FlagValue]string{
	FlagValueNone:  "None",
	FlagValueAllow: "Allow",
	FlagValueDeny:  "Deny",
}

var EnumValuesFlagValue = map[string]FlagValue{
	"None":  FlagValueNone,
	"Allow": FlagValueAllow,
	"Deny":  FlagValueDeny,
}

func (v FlagValue) String() string {
	if s, ok := EnumNamesFlagValue[v]; ok {
		return s
	}
	return "FlagValue(" + strconv.FormatInt(int64(v), 10) + ")"
}
