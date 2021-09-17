package zset

type ZSetLevel struct {
	forward *ZSetNode
	span    int64
}

type ZSetNode struct {
	key      uint64      // unique key of this node
	Value    interface{} // associated data
	score    uint64       // score to determine the order of this node in the set
	previous *ZSetNode
	level    []ZSetLevel
}

func (this *ZSetNode) Key() uint64 {
	return this.key
}

func (this *ZSetNode) Score() uint64 {
	return this.score
}
