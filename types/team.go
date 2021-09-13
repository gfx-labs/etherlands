package types

import (
	"sync"

  flatbuffers "github.com/google/flatbuffers/go"
  proto "github.com/gfx-labs/etherlands/proto"
)

type Group struct {
  name string
  team *Team
  members []*Gamer

  sync.RWMutex
}

func (G *Group) Name() string {
  G.RLock();
  defer G.RUnlock();
  return G.name;
}

func (G *Group) Team() *Team {
  G.RLock();
  defer G.RUnlock();
  return G.team;
}

func (G *Group) Members() []*Gamer {
  G.RLock();
  defer G.RUnlock();
  return G.members;
}

type Team struct {

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

func (T *Team) DefaultPlayerPermissionMap() PlayerPermissionMap {
  T.RLock()
  defer T.RUnlock()
  return T.defaultPlayerPermissions
}

func (T *Team) DefaultGroupPermissionMap() GroupPermissionMap {
  T.RLock()
  defer T.RUnlock()
  return T.defaultGroupPermissions
}

func (T *Team) Owner() *Gamer{
  T.RLock();
  defer T.RUnlock();
  return T.owner
}

func (T *Team) Name() string {
  T.RLock();
  defer T.RUnlock();
  return T.name
}


func (T *Team) Members() []*Gamer{
  T.RLock();
  defer T.RUnlock();
  return T.members
}

func (T *Team) Managers() []*Gamer{
  T.RLock();
  defer T.RUnlock();
  return T.members;
}


func (T *Team) Groups() []*Group{
  T.RLock();
  defer T.RUnlock();
  return T.Groups();
}

func (T *Team) Districts() []*District{
  T.RLock();
  defer T.RUnlock();
  return T.districts;
}


func BuildTeamGroupPermissionVector(builder *flatbuffers.Builder, target GroupPermissionMap) flatbuffers.UOffsetT{
  gp_o := BuildGroupPermissions(builder,  target)
  proto.TeamStartDefaultGroupPermissionsVector(builder,len(gp_o))
  for _, v := range gp_o {
    builder.PrependUOffsetT(v)
  }
  return builder.EndVector(len(gp_o))
}



func BuildTeamPlayerPermissionVector(builder *flatbuffers.Builder, target PlayerPermissionMap) flatbuffers.UOffsetT{
  pp_o := BuildPlayerPermissions(builder, target)
  proto.TeamStartDefaultPlayerPermissionsVector(builder,len(pp_o))
  for _, v := range pp_o {
    builder.PrependUOffsetT(v)
  }
  return builder.EndVector(len(pp_o))
}
