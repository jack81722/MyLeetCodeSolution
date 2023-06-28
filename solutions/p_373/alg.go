package main

import (
	"container/heap"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	maxHeap := make(MaxPairHeap, 0, k)
	heap.Init(&maxHeap)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			if maxHeap.Len() == k && maxHeap[0][0]+maxHeap[0][1] <= n1+n2 {
				break
			}
			heap.Push(&maxHeap, []int{n1, n2})
			if maxHeap.Len() > k {
				heap.Pop(&maxHeap)
			}
		}
	}
	result := make([][]int, 0, k)
	for maxHeap.Len() > 0 && k > 0 {
		res := heap.Pop(&maxHeap).([]int)
		result = append(result, res)
		k--
	}
	return reverse(result)
}

func reverse(nums [][]int) [][]int {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
	return nums
}

type MaxPairHeap [][]int

func (h MaxPairHeap) Len() int           { return len(h) }
func (h MaxPairHeap) Less(i, j int) bool { return h[i][0]+h[i][1] > h[j][0]+h[j][1] }
func (h MaxPairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxPairHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MaxPairHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
