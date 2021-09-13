package types

import (
	"sync"

	"github.com/google/uuid"
)


type Gamer struct{
  nickname string
  address string
  minecraftId uuid.UUID

  mutex sync.RWMutex
}

func (G *Gamer) Nickname() string {
  G.mutex.RLock();
  defer G.mutex.RUnlock();
  return G.nickname
}

func (G *Gamer) Address() string {
  G.mutex.RLock();
  defer G.mutex.RUnlock();
  return G.address
}

func (G *Gamer) MinecraftId() uuid.UUID {
  G.mutex.RLock();
  defer G.mutex.RUnlock();
  return G.minecraftId;
}
