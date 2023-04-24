package main

import (
	"container/heap"
)

func lastStoneWeight(stones []int) int {
	h := IntHeap(stones)
	heap.Init(&h)
	var x, y, result int
	for h.Len() > 1 {
		x = heap.Pop(&h).(int)
		y = heap.Pop(&h).(int)
		result = abs(x - y)
		if result != 0 {
			heap.Push(&h, result)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return h.Pop().(int)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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
