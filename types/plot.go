package types

import "sync"

type Plot struct {
  x       int64
  z       int64

  plot_id uint64
  district_id uint64
  mutex sync.RWMutex
}

func (P *Plot) X() int64{
  return P.x
}

func (P *Plot) Z() int64{
  return P.z
}

func (P *Plot) PlotId() uint64 {
  return P.plot_id
}

func (P *Plot) DistrictId() uint64 {
  P.mutex.RLock()
  defer P.mutex.RUnlock();
  return P.district_id
}

func (P *Plot) SetDistrictId(id uint64){
  P.mutex.Lock()
  defer P.mutex.Unlock()
  P.district_id = id
}

func NewPlot(x, z int64, plotId, districtId uint64) (*Plot) {
  return &Plot{
    plot_id:plotId,
    district_id:districtId,
    x:x,
    z:z,
  }
}
