package types

import (
	"sync"

  flatbuffers "github.com/google/flatbuffers/go"
  proto "github.com/gfx-labs/etherlands/proto"
)

type Group struct {
  name string
  town *Town
  members []*Gamer

  sync.RWMutex
}

func (G *Group) Name() string {
  G.RLock();
  defer G.RUnlock();
  return G.name;
}

func (G *Group) Town() *Town {
  G.RLock();
  defer G.RUnlock();
  return G.town;
}

func (G *Group) Members() []*Gamer {
  G.RLock();
  defer G.RUnlock();
  return G.members;
}

type Town struct {

  name string

  owner *Gamer
  members []*Gamer
  managers []*Gamer

  groups map[string][]*Group
  districts []*District

  defaultPlayerPermissions PlayerPermissionMap
  defaultGroupPermissions GroupPermissionMap

  sync.RWMutex
}

func (T *Town) DefaultPlayerPermissionMap() PlayerPermissionMap {
  T.RLock()
  defer T.RUnlock()
  return T.defaultPlayerPermissions
}

func (T *Town) DefaultGroupPermissionMap() GroupPermissionMap {
  T.RLock()
  defer T.RUnlock()
  return T.defaultGroupPermissions
}

func (T *Town) Owner() *Gamer{
  T.RLock();
  defer T.RUnlock();
  return T.owner
}

func (T *Town) Name() string {
  T.RLock();
  defer T.RUnlock();
  return T.name
}


func (T *Town) Members() []*Gamer{
  T.RLock();
  defer T.RUnlock();
  return T.members
}

func (T *Town) Managers() []*Gamer{
  T.RLock();
  defer T.RUnlock();
  return T.members;
}


func (T *Town) Groups() []*Group{
  T.RLock();
  defer T.RUnlock();
  return T.Groups();
}

func (T *Town) Districts() []*District{
  T.RLock();
  defer T.RUnlock();
  return T.districts;
}


func BuildTownGroupPermissionVector(builder *flatbuffers.Builder, target GroupPermissionMap) flatbuffers.UOffsetT{
  gp_o := BuildGroupPermissions(builder,  target)
  proto.TownStartDefaultGroupPermissionsVector(builder,len(gp_o))
  for _, v := range gp_o {
    builder.PrependUOffsetT(v)
  }
  return builder.EndVector(len(gp_o))
}



func BuildTownPlayerPermissionVector(builder *flatbuffers.Builder, target PlayerPermissionMap) flatbuffers.UOffsetT{
  pp_o := BuildPlayerPermissions(builder, target)
  proto.TownStartDefaultPlayerPermissionsVector(builder,len(pp_o))
  for _, v := range pp_o {
    builder.PrependUOffsetT(v)
  }
  return builder.EndVector(len(pp_o))
}
