package main

import (
	"sort"
)

type KthLargest struct {
	k    int
	nums []int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	return KthLargest{
		k:    k,
		nums: nums,
	}
}

func (this *KthLargest) Add(val int) int {
	for i := range this.nums {
		if this.nums[i] < val {
			this.nums = append(this.nums[:i+1], this.nums[i:]...)
			this.nums[i] = val
			return this.nums[this.k-1]
		}
	}
	this.nums = append(this.nums, val)
	return this.nums[this.k-1]
}
