package main

type uInt64Set map[uint64]struct{}

func (set uInt64Set) add(item uint64) {
	set[item] = struct{}{}
}

func (set uInt64Set) remove(item uint64) {
	delete(set, item)
}
