package types

type Plot struct {
  chainId uint64
  x       int64
  z       int64
}

func (P *Plot) X() int64{
  return P.x
}


func (P *Plot) Z() int64{
  return P.z
}

func (P *Plot) ChainId() uint64 {
  return P.chainId
}

func NewPlot(x, z int64, chainId uint64) (*Plot) {
  return &Plot{
    chainId: chainId,
    x:x,
    z:z,
  }
}
