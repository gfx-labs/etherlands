package types

import (
	"sync"

	flatbuffers "github.com/google/flatbuffers/go"
  proto "github.com/gfx-labs/etherlands/proto"
)

type District struct{
  district_id uint64
  owner *Gamer
  owner_address string

  nickname string
  plots []*Plot


  playerPermissions PlayerPermissionMap
  groupPermissions GroupPermissionMap

  mutex sync.RWMutex
}

func NewDistrict(id uint64, ownerAddress string) *District {
  return &District{
    district_id: id,
    owner_address: ownerAddress,
  }
}

func (D *District) DistrictId() uint64 {
  return D.district_id
}

func (D *District) Nickname() string {
  D.mutex.RLock()
  defer D.mutex.RUnlock()

  return D.nickname
}
func (D *District) OwnerAddress() (string){
  D.mutex.RLock()
  defer D.mutex.RUnlock()
  return D.owner_address
}

func (D *District) Owner() (*Gamer){
  D.mutex.RLock()
  defer D.mutex.RUnlock()
  return D.owner
}

func (D *District) Plots() []*Plot  {
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
