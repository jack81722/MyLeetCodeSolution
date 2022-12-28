package main

import (
	"sort"
)

func minStoneSum(piles []int, k int) int {
	sort.Slice(piles, func(i, j int) bool {
		return piles[i] > piles[j]
	})
	for k > 0 {
		val := first(piles)
		setFirst(val-val/2, piles)
		k--
	}
	return sum(piles...)
}

func heapify(index int, heap []int) {
	if index >= len(heap) {
		return
	}
	l := index*2 + 1
	r := index*2 + 2
	if l < len(heap) {
		heapify(l, heap)
	}
	if r < len(heap) {
		heapify(r, heap)
	}
	if l < len(heap) && heap[index] < heap[l] {
		swap(l, index, heap)
	}
	if r < len(heap) && heap[index] < heap[r] {
		swap(r, index, heap)
	}
}

func first(heap []int) int {
	return heap[0]
}

func setFirst(val int, heap []int) {
	heap[0] = val
	// heapify(0, heap)
	refresh(0, heap)
}

func refresh(index int, heap []int) {
	l := index*2 + 1
	r := index*2 + 2
	val, lval, rval := heap[index], -1, -1
	if l < len(heap) {
		lval = heap[l]
	} else {
		return
	}
	if r < len(heap) {
		rval = heap[r]
	}
	if val < rval && rval > lval {
		swap(r, index, heap)
		refresh(r, heap)
		return
	}
	if val < lval {
		swap(l, index, heap)
		refresh(l, heap)
	}
}

func swap(i, j int, arr []int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func sum(arr ...int) int {
	s := arr[0]
	for i := 1; i < len(arr); i++ {
		s += arr[i]
	}
	return s
}
