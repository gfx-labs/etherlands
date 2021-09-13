package types

import (
	"sync"

	flatbuffers "github.com/google/flatbuffers/go"
  proto "github.com/gfx-labs/etherlands/proto"
)

type District struct{
  chainId uint64

  nickname string
  plots []*Plot

  owner *Gamer

  playerPermissions PlayerPermissionMap
  groupPermissions GroupPermissionMap

  mutex sync.RWMutex
}

func (D *District) ChainId() uint64 {
  return D.chainId
}

func (D *District) Nickname() string {
  D.mutex.RLock()
  defer D.mutex.Unlock()
  return D.nickname
}

func (D *District) Owner() (*Gamer){
  D.mutex.RLock()
  defer D.mutex.Unlock()
  return D.owner
}

func (D *District) Plots() []*Plot  {
  D.mutex.RLock()
  defer D.mutex.Unlock()
  return D.plots
}

func (D *District) PlayerPermissions() PlayerPermissionMap{
  D.mutex.RLock()
  defer D.mutex.Unlock()
  return D.playerPermissions
}

func (D *District) GroupPermissions() GroupPermissionMap{
  D.mutex.RLock()
  defer D.mutex.Unlock()
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
