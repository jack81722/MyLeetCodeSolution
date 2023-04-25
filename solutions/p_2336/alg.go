package main

import "container/heap"

type SmallestInfiniteSet struct {
	cur int
	mp  map[int]bool
	h   IntHeap
}

func Constructor() SmallestInfiniteSet {
	set := SmallestInfiniteSet{
		cur: 1,
		mp:  make(map[int]bool),
		h:   IntHeap{},
	}
	return set
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	var v int
	if len(this.h) > 0 {
		v = heap.Pop(&this.h).(int)
		delete(this.mp, v)
		return v
	}

	v = this.cur
	this.cur++
	return v
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.cur || this.mp[num] {
		return
	}
	this.mp[num] = true
	heap.Push(&this.h, num)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
