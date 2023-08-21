package main

import "container/heap"

func findKthLargest(nums []int, k int) int {
	h := make(IntHeap, 0, 1000)
	heap.Init(&h)
	for _, n := range nums {
		heap.Push(&h, n)
	}
	var result int
	for i := 0; i < k; i++ {
		result = heap.Pop(&h).(int)
	}
	return result
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
