package zset

import (
	"math/rand"
)

const SKIPLIST_MAXLEVEL = 32
const SKIPLIST_P = 0.25

type ZSet struct {
	header *ZSetNode
	tail   *ZSetNode
	length int64
	level  int
	dict   map[uint64]*ZSetNode
}

func createNode(level int, score uint64, key uint64, value interface{}) *ZSetNode {
	node := ZSetNode{
		score: score,
		key:   key,
		Value: value,
		level: make([]ZSetLevel, level),
	}
	return &node
}

func randomLevel() int {
	level := 1
	for float64(rand.Int31()&0xFFFF) < float64(SKIPLIST_P*0xFFFF) {
		level += 1
	}
	if level < SKIPLIST_MAXLEVEL {
		return level
	}

	return SKIPLIST_MAXLEVEL
}

func (Z *ZSet) insertNode(score uint64, key uint64, value interface{}) *ZSetNode {
	var update [SKIPLIST_MAXLEVEL]*ZSetNode
	var rank [SKIPLIST_MAXLEVEL]int64

	x := Z.header
	for i := Z.level - 1; i >= 0; i-- {
		if Z.level-1 == i {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}

		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score &&
					x.level[i].forward.key < key)) {
			rank[i] += x.level[i].span
			x = x.level[i].forward
		}
		update[i] = x
	}

	level := randomLevel()

	if level > Z.level {
		for i := Z.level; i < level; i++ {
			rank[i] = 0
			update[i] = Z.header
			update[i].level[i].span = Z.length
		}
		Z.level = level
	}

	x = createNode(level, score, key, value)
	for i := 0; i < level; i++ {
		x.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = x
		x.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}

	for i := level; i < Z.level; i++ {
		update[i].level[i].span = update[i].level[i].span + 1
	}

	if update[0] == Z.header {
		x.previous = nil
	} else {
		x.previous = update[0]
	}
	if x.level[0].forward != nil {
		x.level[0].forward.previous = x
	} else {
		Z.tail = x
	}
	Z.length = Z.length + 1
	return x
}

func (Z *ZSet) deleteNode(x *ZSetNode, update [SKIPLIST_MAXLEVEL]*ZSetNode) {
	for i := 0; i < Z.level; i++ {
		if update[i].level[i].forward == x {
			update[i].level[i].span += x.level[i].span - 1
			update[i].level[i].forward = x.level[i].forward
		} else {
			update[i].level[i].span -= 1
		}
	}
	if x.level[0].forward != nil {
		x.level[0].forward.previous = x.previous
	} else {
		Z.tail = x.previous
	}
	for Z.level > 1 && Z.header.level[Z.level-1].forward == nil {
		Z.level--
	}
	Z.length--
	delete(Z.dict, x.key)
}

func (Z *ZSet) delete(score uint64, key uint64) bool {
	var update [SKIPLIST_MAXLEVEL]*ZSetNode
	x := Z.header
	for i := Z.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			(x.level[i].forward.score < score ||
				(x.level[i].forward.score == score &&
					x.level[i].forward.key < key)) {
			x = x.level[i].forward
		}
		update[i] = x
	}
	x = x.level[0].forward
	if x != nil && score == x.score && x.key == key {
		Z.deleteNode(x, update)
		return true
	}
	return false
}

func CreateZSet() *ZSet {
	sortedSet := ZSet{
		level: 1,
		dict:  make(map[uint64]*ZSetNode),
	}
	sortedSet.header = createNode(SKIPLIST_MAXLEVEL, 0, 0, nil)
	return &sortedSet
}

func (Z *ZSet) GetCount() int {
	return int(Z.length)
}

func (Z *ZSet) AddOrUpdate(key uint64, score uint64, value interface{}) bool {
	var newNode *ZSetNode = nil

	found := Z.dict[key]
	if found != nil {
		if found.score == score {
			found.Value = value
		} else {
			Z.delete(found.score, found.key)
			newNode = Z.insertNode(score, key, value)
		}
	} else {
		newNode = Z.insertNode(score, key, value)
	}

	if newNode != nil {
		Z.dict[key] = newNode
	}
	return found == nil
}

func (Z *ZSet) Remove(key uint64) *ZSetNode {
	found := Z.dict[key]
	if found != nil {
		Z.delete(found.score, found.key)
		return found
	}
	return nil
}

func (Z *ZSet) GetKeysByScore(score uint64) []uint64 {
	var limit int = int((^uint(0)) >> 1)
	var keys []uint64
	if Z.length == 0 {
		return keys
	}
	x := Z.header
	for i := Z.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
		x.level[i].forward.score < score {
			x = x.level[i].forward
		}
	}

	x = x.level[0].forward

	for x != nil && limit > 0 {
		if x.score > score {
			break
		}

		next := x.level[0].forward

		keys = append(keys, x.key)
		limit = limit - 1

		x = next
	}

	return keys
}

func (Z *ZSet) sanitizeIndexes(start int, end int) (int, int, bool) {
	if start < 0 {
		start = int(Z.length) + start + 1
	}
	if end < 0 {
		end = int(Z.length) + end + 1
	}
	if start <= 0 {
		start = 1
	}
	if end <= 0 {
		end = 1
	}

	reverse := start > end
	if reverse {
		start, end = end, start
	}
	return start, end, reverse
}

func (Z *ZSet) findNodeByRank(start int, remove bool) (traversed int, x *ZSetNode, update [SKIPLIST_MAXLEVEL]*ZSetNode) {
	x = Z.header
	for i := Z.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil &&
			traversed+int(x.level[i].span) < start {
			traversed += int(x.level[i].span)
			x = x.level[i].forward
		}
		if remove {
			update[i] = x
		} else {
			if traversed+1 == start {
				break
			}
		}
	}
	return
}

func (Z *ZSet) GetByRankRange(start int, end int, remove bool) []*ZSetNode {
	start, end, reverse := Z.sanitizeIndexes(start, end)

	var nodes []*ZSetNode

	traversed, x, update := Z.findNodeByRank(start, remove)

	traversed = traversed + 1
	x = x.level[0].forward
	for x != nil && traversed <= end {
		next := x.level[0].forward

		nodes = append(nodes, x)

		if remove {
			Z.deleteNode(x, update)
		}

		traversed = traversed + 1
		x = next
	}

	if reverse {
		for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
			nodes[i], nodes[j] = nodes[j], nodes[i]
		}
	}
	return nodes
}

func (Z *ZSet) GetByRank(rank int, remove bool) *ZSetNode {
	nodes := Z.GetByRankRange(rank, rank, remove)
	if len(nodes) == 1 {
		return nodes[0]
	}
	return nil
}

func (Z *ZSet) GetByKey(key uint64) *ZSetNode {
	return Z.dict[key]
}

func (Z *ZSet) FindRank(key uint64) int {
	var rank int = 0
	node := Z.dict[key]
	if node != nil {
		x := Z.header
		for i := Z.level - 1; i >= 0; i-- {
			for x.level[i].forward != nil &&
				(x.level[i].forward.score < node.score ||
					(x.level[i].forward.score == node.score &&
						x.level[i].forward.key <= node.key)) {
				rank += int(x.level[i].span)
				x = x.level[i].forward
			}

			if x.key == key {
				return rank
			}
		}
	}
	return 0
}

func (Z *ZSet) IterFuncByRankRange(start int, end int, fn func(key uint64, value interface{}) bool) {
	if fn == nil {
		return
	}

	start, end, reverse := Z.sanitizeIndexes(start, end)
	traversed, x, _ := Z.findNodeByRank(start, false)
	var nodes []*ZSetNode

	x = x.level[0].forward
	for x != nil && traversed < end {
		next := x.level[0].forward

		if reverse {
			nodes = append(nodes, x)
		} else if !fn(x.key, x.Value) {
			return
		}

		traversed = traversed + 1
		x = next
	}

	if reverse {
		for i := len(nodes) - 1; i >= 0; i-- {
			if !fn(nodes[i].key, nodes[i].Value) {
				return
			}
		}
	}
}
