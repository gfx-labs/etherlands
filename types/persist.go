package types

import (
	"encoding/binary"
	"errors"
	"fmt"
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


func LoadPlot(chain_id uint64) (*Plot, error){
  bytes, err := ReadStruct("plots", strconv.FormatUint(chain_id,10))
  if err != nil {
    return nil, err
  }
  if len(bytes) < 8 {
    return nil, errors.New(fmt.Sprintf("Empty file for %d",chain_id))
  }
  read_plot := proto.GetRootAsPlot(bytes, 0)
  return NewPlot(
    read_plot.X(),
    read_plot.Z(),
    read_plot.PlotId(),
    read_plot.DistrictId(),
  ), nil
}
func (P *Plot) Save() error {
  builder := flatbuffers.NewBuilder(1024)
  proto.PlotStart(builder)
  proto.PlotAddPlotId(builder, P.PlotId())
  proto.PlotAddDistrictId(builder, P.DistrictId())
  proto.PlotAddX(builder, P.X())
  proto.PlotAddZ(builder, P.Z())
  plot := proto.PlotEnd(builder)
  builder.Finish(plot)

  buf := builder.FinishedBytes()
  return WriteStruct("plots",strconv.FormatUint(P.PlotId(), 10),buf)
}

func LoadDistrict(chain_id uint64) (*District, error){
  bytes, err := ReadStruct("districts", strconv.FormatUint(chain_id,10))
  if err != nil {
    return nil, err
  }
  if len(bytes) < 8 {
    return nil, errors.New(fmt.Sprintf("Empty file for %d",chain_id))
  }
  read_district := proto.GetRootAsDistrict(bytes, 0)
  slice_name := read_district.Nickname()
  fixed_name := [24]byte{}
  for i:=0; (i < len(slice_name)) && (i < 24); i++{
    if(i >= len(slice_name)){
      fixed_name[i] = 0;
    }else{
      fixed_name[i] = slice_name[i];
    }
  }

  return NewDistrict(
    read_district.ChainId(),
    string(read_district.OwnerAddress()),
    fixed_name,
  ), nil
}


func (D *District) Save() error {
  builder := flatbuffers.NewBuilder(1024)

  nickname_offset :=builder.CreateByteVector((D.Nickname())[:])
  proto.DistrictStartPlotsVector(builder, len(D.Plots()))
  for _, v := range D.Plots() {
    builder.PrependUint64(v.PlotId())
  }
  plots_offset := builder.EndVector(len(D.Plots()))
  owner_address_offset :=builder.CreateString(D.OwnerAddress())

  player_permission_offset := BuildDistrictPlayerPermissionVector(builder,D.PlayerPermissions())
  group_permission_offset := BuildDistrictGroupPermissionVector(builder,D.GroupPermissions())

  proto.DistrictStart(builder)

  proto.DistrictAddChainId(builder, D.DistrictId())

  proto.DistrictAddNickname(builder,nickname_offset)

  if(D.Owner() != nil){
    owner_uuid_offset := BuildUUID(builder,D.Owner().MinecraftId())
    proto.DistrictAddOwnerUuid(builder, owner_uuid_offset)
  }
  proto.DistrictAddOwnerAddress(builder, owner_address_offset)
  proto.DistrictAddPlots(builder, plots_offset)
  proto.DistrictAddGroupPermissions(builder, group_permission_offset)
  proto.DistrictAddPlayerPermissions(builder, player_permission_offset)

  //finish
  district_offset := proto.DistrictEnd(builder)
  builder.Finish(district_offset)
  buf := builder.FinishedBytes()

  return WriteStruct("districts",strconv.FormatUint(D.DistrictId(),10),buf)
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
    builder.PrependUint64(v.DistrictId())
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
