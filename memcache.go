package main

import (
	"context"

	"github.com/mediocregopher/radix/v4"
)

type MemoryCache struct {
	redis radix.Client
	ctx   *context.Context
}

func NewMemoryCache() (*MemoryCache, error) {
	ctx := context.Background()
	redis, err := (radix.PoolConfig{}).New(ctx, "tcp", "127.0.0.1:6379")
	if err != nil {
		return nil, err
	}
	return &MemoryCache{redis: redis, ctx: &ctx}, nil
}

func (M *MemoryCache) CacheBlockNumber(blockNumber uint64) error {
	return M.redis.Do(*M.ctx, radix.FlatCmd(nil, "SET", "reader_last_block", blockNumber))
}

func (M *MemoryCache) GetBlockNumber(bn *uint64) error {
	return M.redis.Do(*M.ctx, radix.Cmd(bn, "GET", "reader_last_block"))
}
