package types

import (
	"io/ioutil"
	"os"
	"path/filepath"

	proto "github.com/gfx-labs/etherlands/proto"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/google/uuid"
)

var filelock = MapLock{}

func WriteStruct(root, file string, data []byte) error {
	path := filepath.Join(".", "db", root)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	filepath := path + "/" + file
	unlock := filelock.lock_str(filepath)
	defer unlock()
	os.Remove(filepath)
	err = ioutil.WriteFile(filepath, data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func DeleteStruct(root, file string) error {
	path := filepath.Join(".", "db", root)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	filepath := path + "/" + file
	unlock := filelock.lock_str(filepath)
	defer unlock()
	err = os.Remove(filepath)
	if err != nil {
		return err
	}
	return nil
}

func ReadStruct(root, file string) ([]byte, error) {
	path := filepath.Join(".", "db", root)
	filepath := path + "/" + file
	unlock := filelock.lock_str(filepath)
	defer unlock()
	return ioutil.ReadFile(filepath)
}

func ListStruct(root string) ([]os.FileInfo, error) {
	path := filepath.Join(".", "db", root)
	return ioutil.ReadDir(path)
}

func BuildUUID(builder *flatbuffers.Builder, gamerId uuid.UUID) flatbuffers.UOffsetT {
	return proto.CreateUUID(builder, gamerId[0],
		gamerId[1],
		gamerId[2],
		gamerId[3],
		gamerId[4],
		gamerId[5],
		gamerId[6],
		gamerId[7],
		gamerId[8],
		gamerId[9],
		gamerId[10],
		gamerId[11],
		gamerId[12],
		gamerId[13],
		gamerId[14],
		gamerId[15],
	)
}

type PlayerPermissionEntry struct {
	uuid  uuid.UUID
	flag  proto.AccessFlag
	value proto.FlagValue
}

func BuildTeamPermissions(
	builder *flatbuffers.Builder,
	target *TeamPermissionMap,
) []flatbuffers.UOffsetT {
	gp_o := []flatbuffers.UOffsetT{}
	for _, v := range FlattenTeamPermissionMap(target) {
		team_name := builder.CreateString(v.name)
		proto.TeamPermissionStart(builder)
		proto.TeamPermissionAddFlag(builder, v.flag)
		proto.TeamPermissionAddValue(builder, v.value)
		proto.TeamPermissionAddTeam(builder, team_name)
		entry := proto.TeamPermissionEnd(builder)
		gp_o = append(gp_o, entry)
	}
	return gp_o
}

func BuildPlayerPermissions(
	builder *flatbuffers.Builder,
	target *PlayerPermissionMap,
) []flatbuffers.UOffsetT {
	pp_o := []flatbuffers.UOffsetT{}
	for _, v := range FlattenPlayerPermissionMap(target) {
		proto.PlayerPermissionStart(builder)
		proto.PlayerPermissionAddFlag(builder, v.flag)
		proto.PlayerPermissionAddValue(builder, v.value)
		player_uuid := BuildUUID(builder, v.uuid)
		proto.PlayerPermissionAddMinecraftId(builder, player_uuid)
		entry := proto.PlayerPermissionEnd(builder)
		pp_o = append(pp_o, entry)
	}
	return pp_o
}

func FlattenPlayerPermissionMap(target *PlayerPermissionMap) []PlayerPermissionEntry {
	output := []PlayerPermissionEntry{}
	for id, map_value := range target.i {
		for flag, value := range map_value {
			output = append(output, PlayerPermissionEntry{
				uuid:  id,
				flag:  flag,
				value: value,
			})
		}
	}
	return output
}

type TeamPermissionEntry struct {
	name  string
	flag  proto.AccessFlag
	value proto.FlagValue
}

func FlattenTeamPermissionMap(target *TeamPermissionMap) []TeamPermissionEntry {
	output := []TeamPermissionEntry{}
	for id, map_value := range target.i {
		for flag, value := range map_value {
			output = append(output, TeamPermissionEntry{
				name:  id,
				flag:  flag,
				value: value,
			})
		}
	}
	return output
}
