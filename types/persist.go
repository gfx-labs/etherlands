package types

import (
  "encoding/binary"
  "strconv"

  proto "github.com/gfx-labs/etherlands/proto"
  flatbuffers "github.com/google/flatbuffers/go"
  "github.com/google/uuid"
)

func BreakUUID(id uuid.UUID) (uint64, uint64) {
  l1 := binary.BigEndian.Uint64(id[:8])
  l2 := binary.BigEndian.Uint64(id[8:])
  return l1, l2
}


func (G *Gamer) Save() error{
  builder:=flatbuffers.NewBuilder(1024);
  proto.GamerStart(builder)

  addr:=builder.CreateString(G.Address());
  proto.GamerAddAddress(builder,addr)

  nick:=builder.CreateString(G.Nickname())
  proto.GamerAddNickname(builder,nick)

  gamer := proto.GamerEnd(builder);
  builder.Finish(gamer);

  buf := builder.FinishedBytes()
  return WriteStruct("gamers",G.MinecraftId().String(),buf)
}

func (D *District) Save() error {
  builder:=flatbuffers.NewBuilder(1024);
  proto.DistrictStart(builder);
  proto.DistrictAddChainId(builder, D.ChainId())
  hi, lo := BreakUUID(D.Owner().MinecraftId())
  owner_id := proto.CreateUUID(builder,hi,lo)
  proto.DistrictAddOwnerUuid(builder, owner_id)
  return nil
}


func LoadPlot(chain_id uint64) (*Plot, error){
  bytes, err := ReadStruct("plots", strconv.FormatUint(chain_id,10))
  if err != nil {
    return nil, err
  }
  read_plot :=proto.GetRootAsPlot(bytes, 0)
  return NewPlot(
    read_plot.X(),
    read_plot.Z(),
    read_plot.ChainId(),
  ), nil

}
func (P *Plot) Save() error {
  builder := flatbuffers.NewBuilder(1024)
  proto.PlotStart(builder)
  proto.PlotAddChainId(builder, P.ChainId())
  proto.PlotAddX(builder, P.X())
  proto.PlotAddZ(builder, P.Z())
  plot := proto.PlotEnd(builder)
  builder.Finish(plot)

  buf := builder.FinishedBytes()
  return WriteStruct("plots",strconv.FormatUint(P.ChainId(), 10),buf)
}

func (D *District) save() error {
  builder := flatbuffers.NewBuilder(1024)

  nickname_offset :=builder.CreateString(D.Nickname())
  proto.DistrictStartPlotsVector(builder, len(D.Plots()))
  for _, v := range D.Plots() {
    builder.PrependUint64(v.ChainId())
  }
  plots_offset := builder.EndVector(len(D.Plots()))
  owner_uuid_offset := BuildUUID(builder,D.Owner().MinecraftId())
  owner_address_offset :=builder.CreateString(D.Nickname())
  player_permission_offset := BuildDistrictPlayerPermissionVector(builder,D.PlayerPermissions())
  group_permission_offset := BuildDistrictGroupPermissionVector(builder,D.GroupPermissions())

  proto.DistrictStart(builder)

  proto.DistrictAddChainId(builder, D.ChainId())

  proto.DistrictAddNickname(builder,nickname_offset)
  proto.DistrictAddOwnerUuid(builder, owner_uuid_offset)
  proto.DistrictAddOwnerAddress(builder, owner_address_offset)
  proto.DistrictAddPlots(builder, plots_offset)
  proto.DistrictAddGroupPermissions(builder, group_permission_offset)
  proto.DistrictAddPlayerPermissions(builder, player_permission_offset)

  return nil;
}

func (T *Team) Save() error {
  builder := flatbuffers.NewBuilder(1024)
  // create default player permission vector
  player_permission_offset := BuildTeamPlayerPermissionVector(builder,T.defaultPlayerPermissions)
  // create default group permission vector
  group_permission_offset := BuildTeamGroupPermissionVector(builder,T.defaultGroupPermissions)

  // create districts vector
  proto.TeamStartDistrictsVector(builder, len(T.Districts()))
  for _, v := range T.Districts() {
    builder.PrependUint64(v.ChainId())
  }
  districts_offset := builder.EndVector(len(T.Districts()))

  // create team manager vector
  team_managers := T.Managers()
  proto.TeamStartManagersVector(builder,len(team_managers))
  for _ , v := range team_managers {
    manager_offset := BuildUUID(builder, v.MinecraftId())
    builder.PrependUOffsetT(manager_offset)
  }
  manager_vector := builder.EndVector(len(team_managers))

  // create team member vector
  team_members := T.Members()
  proto.TeamStartMembersVector(builder,len(team_members))
  for _ , v := range team_members {
    member_offset := BuildUUID(builder, v.MinecraftId())
    builder.PrependUOffsetT(member_offset)
  }
  member_vector := builder.EndVector(len(team_members))


  //create team table
  proto.TeamStart(builder)
  //team name
  team_name := builder.CreateString(T.Name())
  proto.TeamAddName(builder, team_name)

  //owner
  owner_id := BuildUUID(builder,T.Owner().MinecraftId())
  proto.TeamAddOwner(builder, owner_id)

  //members
  proto.TeamAddMembers(builder,member_vector)

  //managers
  proto.TeamAddManagers(builder,manager_vector)
  //districts
  proto.TeamAddDistricts(builder, districts_offset)
  //perms
  proto.TeamAddDefaultGroupPermissions(builder, group_permission_offset);
  proto.TeamAddDefaultPlayerPermissions(builder, player_permission_offset);

  //finish
  team_offset := proto.TeamEnd(builder)
  builder.Finish(team_offset)
  buf := builder.FinishedBytes()


  return WriteStruct("team",T.Name(),buf)
}
