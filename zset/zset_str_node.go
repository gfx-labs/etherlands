package zset

type ZSetStrLevel struct {
	forward *ZSetStrNode
	span    int64
}

type ZSetStrNode struct {
	key      uint64
	Value    interface{}
	score    string
	previous *ZSetStrNode
	level    []ZSetStrLevel
}

func (this *ZSetStrNode) Key() uint64 {
	return this.key
}

func (this *ZSetStrNode) Score() string {
	return this.score
}
