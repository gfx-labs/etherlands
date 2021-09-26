package types

import (
	"errors"
	"fmt"

	"github.com/google/uuid"

  proto "github.com/gfx-labs/etherlands/proto"
)

// note - loading a gamer from memory will pause everything
// design improvement will be to make this a gamer distributor that sends to channels
// but w/e we will do that... another decade....
func (W *World) GetGamer(gamer_id uuid.UUID) *Gamer {
  // first check if the gamer is in live cache
  W.gamers_lock.RLock();
  if val, ok := W.gamers[NewGamerKey(gamer_id)]; ok{
    W.gamers_lock.RUnlock();
    return val
  }
  //release the read lock
  W.gamers_lock.RUnlock();

  //obtain a write lock
  W.gamers_lock.Lock();
  defer W.gamers_lock.Unlock();
  // if not in live cache, see if gamer file exists
  if res, err := W.LoadGamer(gamer_id); err == nil {
    // add it to the cache
    W.gamers[NewGamerKey(gamer_id)] = res
    return res
  }
  // oh no!! the gamer does not exist!!! make one!!!
  output := W.newGamer(gamer_id)
  // add it to the cache
  W.gamers[NewGamerKey(gamer_id)] = output
  return output
}

func (W *World) newGamer(gamer_id uuid.UUID) (*Gamer) {
  return &Gamer{
    world: W,
    minecraftId: gamer_id,
  }
}

func (W *World) LoadGamer(gamer_id uuid.UUID) (*Gamer, error){
  bytes, err := ReadStruct("gamers", gamer_id.String())
  if err != nil {
    return nil, err
  }
  if len(bytes) < 8 {
    return nil, errors.New(fmt.Sprintf("Empty file for %s",gamer_id.String()))
  }
  read_gamer := proto.GetRootAsGamer(bytes, 0)

  read_gam := read_gamer.MinecraftId(nil)
  read_uuid  := [16]byte{
    byte(read_gam.B0()),
    byte(read_gam.B1()),
    byte(read_gam.B2()),
    byte(read_gam.B3()),
    byte(read_gam.B4()),
    byte(read_gam.B5()),
    byte(read_gam.B6()),
    byte(read_gam.B7()),
    byte(read_gam.B8()),
    byte(read_gam.B9()),
    byte(read_gam.B10()),
    byte(read_gam.B11()),
    byte(read_gam.B12()),
    byte(read_gam.B13()),
    byte(read_gam.B14()),
    byte(read_gam.B15()),
  }

  return &Gamer{
    world: W,
    minecraftId:read_uuid,
    address: string(read_gamer.Address()),
    nickname: string(read_gamer.Nickname()),
  }, nil
}
