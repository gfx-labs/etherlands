package types

import (
	"io/ioutil"
	"os"
	"path/filepath"

	proto "github.com/gfx-labs/etherlands/proto"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/google/uuid"
)


func WriteStruct(root, file string, data []byte) error {
  path := filepath.Join(".", "db", root)
  err := os.MkdirAll(path, os.ModePerm)
  if err != nil {
    return err
  }
  filepath := path + "/" + file
  err = ioutil.WriteFile(filepath, data, 0666)
  if err != nil {
    return err
  }
  return nil
}

func ReadStruct(root, file string) ([]byte, error) {
  path := filepath.Join(".", "db", root)
  filepath := path + "/" + file
  return ioutil.ReadFile(filepath)
}

func BuildUUID(builder *flatbuffers.Builder, gamerId uuid.UUID) flatbuffers.UOffsetT {
  return proto.CreateUUID(builder, int8(gamerId[0]),
    int8(gamerId[1]),
    int8(gamerId[2]),
    int8(gamerId[3]),
    int8(gamerId[4]),
    int8(gamerId[5]),
    int8(gamerId[6]),
    int8(gamerId[7]),
    int8(gamerId[8]),
    int8(gamerId[9]),
    int8(gamerId[10]),
    int8(gamerId[11]),
    int8(gamerId[12]),
    int8(gamerId[13]),
    int8(gamerId[14]),
    int8(gamerId[15]),
  )
}


type PlayerPermissionEntry struct {
  uuid uuid.UUID
  flag proto.AccessFlag
  value proto.FlagValue
}

func BuildGroupPermissions(builder *flatbuffers.Builder, target GroupPermissionMap) []flatbuffers.UOffsetT {
  gp_o := []flatbuffers.UOffsetT{}
  for _, v := range  FlattenGroupPermissionMap(target) {
    proto.GroupPermissionStart(builder)
    proto.GroupPermissionAddFlag(builder,v.flag)
    proto.GroupPermissionAddValue(builder,v.value)
    group_name := builder.CreateString(v.name)
    proto.GroupPermissionAddGroup(builder, group_name)
    entry := proto.GroupPermissionEnd(builder)
    gp_o = append(gp_o,entry)
  }
  return gp_o
}


func BuildPlayerPermissions(builder *flatbuffers.Builder, target PlayerPermissionMap) []flatbuffers.UOffsetT {
  pp_o := []flatbuffers.UOffsetT{}
  for _, v := range  FlattenPlayerPermissionMap(target) {
    proto.PlayerPermissionStart(builder)
    proto.PlayerPermissionAddFlag(builder,v.flag)
    proto.PlayerPermissionAddValue(builder,v.value)
    //hi, lo := BreakUUID(v.uuid);
    //player_uuid := proto.CreateUUID(builder,hi,lo)
    //proto.PlayerPermissionAddMinecraftId(builder, player_uuid)
    entry := proto.PlayerPermissionEnd(builder)
    pp_o = append(pp_o,entry)
  }
  return pp_o
}


func FlattenPlayerPermissionMap(target PlayerPermissionMap) ([]PlayerPermissionEntry) {
  output := []PlayerPermissionEntry{}
  for id, map_value := range target {
    for flag, value := range map_value{
      output = append(output, PlayerPermissionEntry{
        uuid:id,
        flag:flag,
        value:value,
      })
    }
  }
  return output
}


type GroupPermissionEntry struct {
  name string
  flag proto.AccessFlag
  value proto.FlagValue
}


func FlattenGroupPermissionMap(target GroupPermissionMap) ([]GroupPermissionEntry) {
  output := []GroupPermissionEntry{}
  for id, map_value := range target {
    for flag, value := range map_value{
      output = append(output, GroupPermissionEntry{
        name:id,
        flag:flag,
        value:value,
      })
    }
  }
  return output
}
