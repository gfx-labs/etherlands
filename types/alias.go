package types

import (
  uuid "github.com/google/uuid"
  proto "github.com/gfx-labs/etherlands/proto"
)

type(
  PlayerPermissionMap = map[uuid.UUID]map[proto.AccessFlag]proto.FlagValue
  GroupPermissionMap map[string]map[proto.AccessFlag]proto.FlagValue
)
