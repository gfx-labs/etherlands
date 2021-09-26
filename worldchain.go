package main

import (
	"context"
	"log"
	"math/big"
	"time"

	//"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	types "github.com/gfx-labs/etherlands/types"
)

const CONTRACT_ADDR = "0x8ed31d7ff5d2ffbf17fe3118a61123f50adb523a"
const RPC_ADDR = "https://polygon-mainnet.g.alchemy.com/v2/mkvOLrm_XUrvBB5emIKb7AimZDOWct6c"
const RPC_MAX = 1999
const FIRST_BLOCK = 19194856

type DistrictConnection struct {
	provider *ethclient.Client
	contract *DistrictContract
	ctx      *context.Context

	cache *MemoryCache

	W *types.World

	best_block uint64
}

func (D *DistrictConnection) GetPlotInfo(plot_id uint64) (types.PlotChainInfo, error) {
	big_id := big.NewInt(int64(plot_id))
	output := types.PlotChainInfo{}
	x, err := D.contract.PlotX(&bind.CallOpts{Pending: false}, big_id)
	if err != nil {
		return output, err
	}
	z, err := D.contract.PlotZ(&bind.CallOpts{Pending: false}, big_id)
	if err != nil {
		return output, err
	}
	district_id, err := D.contract.PlotDistrictOf(&bind.CallOpts{Pending: false}, big_id)
	if err != nil {
		return output, err
	}
	output.X = x.Int64()
	output.Z = z.Int64()
	output.PlotId = plot_id
	output.DistrictId = district_id.Uint64()
	return output, nil
}

func (D *DistrictConnection) GetDistrictInfo(district_id uint64) (types.DistrictChainInfo, error) {
	big_id := big.NewInt(int64(district_id))
	output := types.DistrictChainInfo{}
	x, err := D.contract.OwnerOf(&bind.CallOpts{Pending: false}, big_id)
	if err != nil {
		return output, err
	}
	name_bytes, err := D.contract.DistrictNameOf(&bind.CallOpts{Pending: false}, big_id)
	if err != nil {
		return output, err
	}
	output.DistrictId = district_id
	output.Nickname = name_bytes
	output.Owner = x.String()
	return output, nil
}

func NewDistrictConnection(W *types.World) (*DistrictConnection, error) {
	ctx := context.Background()
	client, err := ethclient.DialContext(ctx, RPC_ADDR)
	if err != nil {
		return nil, err
	}
	contract, err := NewDistrictContract(common.HexToAddress(CONTRACT_ADDR), client)
	if err != nil {
		return nil, err
	}
	output := &DistrictConnection{contract: contract,
		provider:   client,
		ctx:        &ctx,
		best_block: FIRST_BLOCK,
		W:          W,
	}
	memcache, err := NewMemoryCache()
	if err == nil {
		var num uint64
		err = memcache.GetBlockNumber(&num)
		if err == nil {
			output.best_block = num
		}
		output.cache = memcache
	} else {
		log.Println("no redis - running in dumb mode")
	}

	return output, nil
}

func (D *DistrictConnection) GetTotalPlots() (uint64, error) {
	b, err := D.contract.TotalPlots(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}
	return b.Uint64(), err
}

func (D *DistrictConnection) GetTotalDistricts() (uint64, error) {
	b, err := D.contract.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}
	return b.Uint64(), err
}

func (D *DistrictConnection) QueryRecentEvents() (uint64, error) {
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
		End:   &target,
	}, nil, nil, nil)

	if err != nil {
		return D.best_block, err
	}
	plot_transfer_logs, err := D.contract.FilterPlotTransfer(&bind.FilterOpts{
		Start: D.best_block,
		End:   &target,
	})
	if err != nil {
		return D.best_block, err
	}
	plot_creation_logs, err := D.contract.FilterPlotCreation(&bind.FilterOpts{
		Start: D.best_block,
		End:   &target,
	})
	if err != nil {
		return D.best_block, err
	}

	district_name_logs, err := D.contract.FilterDistrictName(&bind.FilterOpts{
		Start: D.best_block,
		End:   &target,
	})
	if err != nil {
		return D.best_block, err
	}

	district_queue := make(uInt64Set)
	plot_queue := make(uInt64Set)

	for plot_transfer_logs.Next() {
		district_queue.add(plot_transfer_logs.Event.TargetId.Uint64())
		plot_queue.add(plot_transfer_logs.Event.PlotId.Uint64())
	}

	for district_name_logs.Next() {
		district_queue.add(district_name_logs.Event.DistrictId.Uint64())
	}

	for transfer_logs.Next() {
		district_queue.add(transfer_logs.Event.TokenId.Uint64())
	}

	for plot_creation_logs.Next() {
		plot_queue.add(plot_creation_logs.Event.PlotId.Uint64())
	}

	for k := range district_queue {
		D.W.DistrictRequests <- k
	}

	for k := range plot_queue {
		D.W.PlotRequests <- k
	}

	D.best_block = target + 1
	return target + 1, nil
}
func (D *DistrictConnection) process_events() {
	district_limiter := NewRateLimiter()
	plot_limiter := NewRateLimiter()
	for {
		select {
		case district_id := <-D.W.DistrictRequests:
			if district_limiter.check(district_id) {
				log.Println("updating district", district_id)
				district, err := D.GetDistrictInfo(district_id)
				if err == nil {
					D.W.DistrictIn <- district
				}
			} else {
				log.Println("skipping district update", district_id)
			}
		case plot_id := <-D.W.PlotRequests:
			if plot_limiter.check(plot_id) {
				log.Println("updating plot ", plot_id)
				plot, err := D.GetPlotInfo(plot_id)
				if err == nil {
					D.W.PlotIn <- plot
				}
			} else {
				log.Println("skipping plot update", plot_id)
			}
		}
	}
}

func (D *DistrictConnection) start_query(duration time.Duration) {
	if duration == 0 {
		return
	}
	query_event_timer := start_repeating(duration)
	for {
		_ = <-query_event_timer
		block, err := D.QueryRecentEvents()
		if err != nil {
			log.Println(err)
		} else {
			if D.cache != nil {
				D.cache.CacheBlockNumber(block)
			}
		}
	}
}
