module github.com/gfx-labs/etherlands/types

go 1.17

replace github.com/gfx-labs/etherlands/zset => ../zset

replace github.com/gfx-labs/etherlands/proto => ../proto

require (
	github.com/gfx-labs/etherlands/proto v0.0.0-00010101000000-000000000000
	github.com/gfx-labs/etherlands/zset v0.0.0-00010101000000-000000000000
	github.com/google/flatbuffers v2.0.0+incompatible
	github.com/google/uuid v1.3.0
	github.com/mediocregopher/radix/v4 v4.0.0-beta.1
)

require github.com/tilinna/clock v1.0.2 // indirect
