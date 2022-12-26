package main

import (
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	maxs := make([]int, len(queries))
	for i, q := range queries {
		maxs[i] = binarySearch(q, nums) + 1
	}
	return maxs
}

func binarySearch(query int, nums []int) int {
	l := 0
	r := len(nums)
	var mid int
	for l < r {
		mid = (l + r) / 2
		if nums[mid] == query {
			return mid
		}
		if nums[mid] < query {
			l = mid + 1
		}
		if nums[mid] > query {
			r = mid
		}
	}
	if nums[mid] > query {
		return mid - 1
	}
	return mid
}
