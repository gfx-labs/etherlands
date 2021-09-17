package types

import (
	"fmt"
	"sync"

	proto "github.com/gfx-labs/etherlands/proto"
	flatbuffers "github.com/google/flatbuffers/go"
)

type District struct{
  district_id uint64
  owner *Gamer
  owner_address string

  nickname *[24]byte
  plots map[uint64]*Plot


  playerPermissions PlayerPermissionMap
  groupPermissions GroupPermissionMap

  mutex sync.RWMutex
}

func NewDistrict(id uint64, ownerAddress string, nickname [24]byte) *District {
  return &District{
    district_id: id,
    owner_address: ownerAddress,
    nickname: &nickname,
  }
}

func (D *District) DistrictId() uint64 {
  return D.district_id
}

func (D *District) Nickname() *[24]byte {
  D.mutex.RLock()
  defer D.mutex.RUnlock()
  return D.nickname
}

func (D *District) StringName() string {
  pending := Parse24Name(*D.Nickname())
  if pending == "" {
    return fmt.Sprintf("#%d",D.district_id)
  }
  return pending;
}

func (D *District) OwnerAddress() (string){
  D.mutex.RLock()
  defer D.mutex.RUnlock()
  return D.owner_address
}

func (D *District) SetOwnerAddress(addr string) {
  D.mutex.Lock()
  defer D.mutex.Unlock()
  D.owner_address = addr
}

func (D *District) Owner() (*Gamer){
  D.mutex.RLock()
  defer D.mutex.RUnlock()
  return D.owner
}

func (D *District) Plots() map[uint64]*Plot  {
  D.mutex.RLock()
  defer D.mutex.RUnlock()
  return D.plots
}

func (D *District) PlayerPermissions() PlayerPermissionMap{
  D.mutex.RLock()
  defer D.mutex.RUnlock()
  return D.playerPermissions
}

func (D *District) GroupPermissions() GroupPermissionMap{
  D.mutex.RLock()
  defer D.mutex.RUnlock()
  return D.groupPermissions
}

func BuildDistrictGroupPermissionVector(builder *flatbuffers.Builder, target GroupPermissionMap) flatbuffers.UOffsetT{
  gp_o := BuildGroupPermissions(builder,  target)
  proto.DistrictStartGroupPermissionsVector(builder,len(gp_o))
  for _, v := range gp_o {
    builder.PrependUOffsetT(v)
  }
  return builder.EndVector(len(gp_o))
}

func BuildDistrictPlayerPermissionVector(builder *flatbuffers.Builder, target PlayerPermissionMap) flatbuffers.UOffsetT{
  pp_o := BuildPlayerPermissions(builder, target)
  proto.DistrictStartPlayerPermissionsVector(builder,len(pp_o))
  for _, v := range pp_o {
    builder.PrependUOffsetT(v)
  }
  return builder.EndVector(len(pp_o))
}
