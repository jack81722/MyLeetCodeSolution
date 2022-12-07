package main

import "sort"

func maximumProduct(nums []int, k int) int {
	mod := 1000000007
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	min := nums[0]
	i := 0
	for k > 0 {
		if i >= len(nums) {
			i = 0
		}
		n := nums[i]
		if n > min {
			min = min + 1
			i = 0
		}
		nums[i]++
		k--
		i++
	}
	result := 1
	for _, n := range nums {
		result = (result * n) % mod
	}
	return result
}
