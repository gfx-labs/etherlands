package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gfx-labs/etherlands/types"
)

const CONTRACT_ADDR = "0xc7B4Cdf2c8ff3FC94D4f9f882D86CE824e0FB985"

type DistrictConnection struct {
	provider *ethclient.Client
	contract *DistrictContract
	ctx *context.Context
}



func NewDistrictConnection() (*DistrictConnection,error) {
	ctx := context.Background();
	client, err := ethclient.DialContext(ctx, "https://polygon-rpc.com")
	if err != nil {
		return nil, err
	}
	contract, err := NewDistrictContract(common.HexToAddress(CONTRACT_ADDR),client)
	if err != nil {
		return nil, err
	}

	return &DistrictConnection{contract: contract, provider:client, ctx:&ctx}, nil
}

func (D *DistrictConnection) GetPlotInfo(plot_id uint64) (*types.Plot, error) {
	big_id := big.NewInt(int64(plot_id))
	x, err := D.contract.PlotX(&bind.CallOpts{Pending:false},big_id)
	if err != nil {
		return nil, err
	}
	z, err := D.contract.PlotZ(&bind.CallOpts{Pending:false},big_id)
	if err != nil {
		return nil, err
	}
	if(z.Int64() == 0 && x.Int64() == 0){
		return nil, errors.New(fmt.Sprintf("Plot %d does not yet exist",plot_id))
	}

	return types.NewPlot(x.Int64(),z.Int64(),plot_id), nil

}

func (D *DistrictConnection) GetTotalPlots() (uint64, error) {
	b, err:= D.contract.TotalPlots(&bind.CallOpts{})
	if(err != nil){
		return 0, err
	}
	return b.Uint64(), err
}
