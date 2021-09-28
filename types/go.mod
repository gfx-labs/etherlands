module github.com/gfx-labs/etherlands/types

go 1.17

replace github.com/gfx-labs/etherlands/zset => ../zset

replace github.com/gfx-labs/etherlands/proto => ../proto

replace github.com/gfx-labs/etherlands/utils => ../utils

require (
	github.com/gfx-labs/etherlands/proto v0.0.0-00010101000000-000000000000
	github.com/gfx-labs/etherlands/utils v0.0.0-00010101000000-000000000000
	github.com/gfx-labs/etherlands/zset v0.0.0-00010101000000-000000000000
	github.com/google/flatbuffers v2.0.0+incompatible
	github.com/google/uuid v1.3.0
)
