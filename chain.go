package main

import (
	"context"
	"errors"
	"fmt"
	//"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gfx-labs/etherlands/types"
)

const CONTRACT_ADDR = "0x8ed31d7ff5d2ffbf17fe3118a61123f50adb523a"
const RPC_ADDR = "https://polygon-mainnet.g.alchemy.com/v2/mkvOLrm_XUrvBB5emIKb7AimZDOWct6c"
const RPC_MAX = 1999
const FIRST_BLOCK = 19194856

type TransferEvent struct {
	to string

	district_id uint64
}

type PlotTransferEvent struct {
	origin_district  uint64
	target_district  uint64

	plot_id uint64
}

type PlotCreationEvent struct{
	x_coord int64
	z_coord int64

	plot_id uint64
}

type DistrictNameEvent struct{
	district_id uint64
}

type DistrictConnection struct {
	provider *ethclient.Client
	contract *DistrictContract
	ctx *context.Context


	PlotChannel chan uint64
	DistrictChannel chan uint64

	best_block uint64
}


func NewDistrictConnection() (*DistrictConnection,error) {
	ctx := context.Background();
	client, err := ethclient.DialContext(ctx, RPC_ADDR)
	if err != nil {
		return nil, err
	}
	contract, err := NewDistrictContract(common.HexToAddress(CONTRACT_ADDR),client)
	if err != nil {
		return nil, err
	}

	return &DistrictConnection{contract: contract,
	provider:client,
	ctx:&ctx,
	PlotChannel: make(chan uint64, 100),
	DistrictChannel: make(chan uint64, 100),
	best_block: FIRST_BLOCK,
}, nil
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
	district_id, err := D.contract.PlotDistrictOf(&bind.CallOpts{Pending:false},big_id)
	if err != nil {
		return nil, err
	}
	if(z.Int64() == 0 && x.Int64() == 0 && plot_id != 25){
		return nil, errors.New(fmt.Sprintf("Plot %d does not yet exist",plot_id))
	}
	return types.NewPlot(x.Int64(),z.Int64(),plot_id, district_id.Uint64()), nil
}

func (D *DistrictConnection) UpdatePlotDistrict(plot *types.Plot) (error) {
	big_id := big.NewInt(int64(plot.PlotId()))
	district_id, err := D.contract.PlotDistrictOf(&bind.CallOpts{Pending:false},big_id)
	if err != nil {
		return err
	}else{
		plot.SetDistrictId(district_id.Uint64())
		return nil
	}
}

func (D *DistrictConnection) GetDistrictInfo(district_id uint64) (*types.District, error) {
	big_id := big.NewInt(int64(district_id))
	x, err := D.contract.OwnerOf(&bind.CallOpts{Pending:false},big_id)
	if err != nil {
		return nil, err
	}
	name_bytes, err := D.contract.DistrictNameOf(&bind.CallOpts{Pending:false},big_id)
	if err != nil {
		return nil, err
	}
	return types.NewDistrict(district_id, x.String(),name_bytes), nil
}

func (D *DistrictConnection) UpdateDistrictOwner(district *types.District) (error) {
	big_id := big.NewInt(int64(district.DistrictId()))
	district_owner, err := D.contract.OwnerOf(&bind.CallOpts{Pending:false},big_id)
	if err != nil {
		return err
	}else{
		district.SetOwnerAddress(district_owner.String())
		return nil
	}
}

func (D *DistrictConnection) GetTotalPlots() (uint64, error) {
	b, err:= D.contract.TotalPlots(&bind.CallOpts{})
	if(err != nil){
		return 0, err
	}
	return b.Uint64(), err
}

func (D *DistrictConnection) GetTotalDistricts() (uint64, error) {
	b, err:= D.contract.TotalSupply(&bind.CallOpts{})
	if(err != nil){
		return 0, err
	}
	return b.Uint64(), err
}


func (D *DistrictConnection) QueryRecentEvents() (uint64,error) {
	current_block, err := D.provider.BlockNumber(*D.ctx)
	if err != nil {
		return D.best_block, err
	}
	target := D.best_block + RPC_MAX
	if current_block < target {
		target = current_block
	}
	transfer_logs, err := D.contract.FilterTransfer(&bind.FilterOpts{
		Start: D.best_block,
		End: &target,
	},nil,nil,nil)

	if err != nil {
		return D.best_block, err
	}
	plot_transfer_logs, err := D.contract.FilterPlotTransfer(&bind.FilterOpts{
		Start: D.best_block,
		End: &target,
	})
	if err != nil {
		return D.best_block, err
	}
	plot_creation_logs, err := D.contract.FilterPlotCreation(&bind.FilterOpts{
		Start: D.best_block,
		End: &target,
	})
	if err != nil {
		return D.best_block, err
	}

	district_name_logs, err := D.contract.FilterDistrictName(&bind.FilterOpts{
		Start: D.best_block,
		End: &target,
	})
	if err != nil {
		return D.best_block, err
	}

	for plot_transfer_logs.Next() {
		D.DistrictChannel <- plot_transfer_logs.Event.TargetId.Uint64()
			D.PlotChannel <- plot_transfer_logs.Event.PlotId.Uint64()
	}

	for district_name_logs.Next() {
			D.DistrictChannel <- district_name_logs.Event.DistrictId.Uint64()
	}

	for transfer_logs.Next() {
			D.DistrictChannel <- transfer_logs.Event.TokenId.Uint64()
	}

	for plot_creation_logs.Next() {
			D.PlotChannel <- plot_creation_logs.Event.PlotId.Uint64()
	}

	D.best_block = target + 1
	return target + 1, nil
}
