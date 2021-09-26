package types


import (
)

type FamilyType = uint64;

const(
  NO_FAMILY FamilyType = iota
  PLOT_FAMILY
  DISTRICT_FAMILY
  GAMER_FAMILY
  TOWN_FAMILY
)

type FamilyKey struct {
  datatype FamilyType
  subkey string
}

type FamilyMember interface {
  GetParentKeys() []FamilyKey
  GetKey() FamilyKey
}
