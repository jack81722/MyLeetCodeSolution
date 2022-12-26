package main

import (
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	sums := make([]int, len(nums))
	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}
	maxs := make([]int, len(queries))
	for i, q := range queries {
		maxs[i] = searchAndCount(0, q, nums, sums)
		for j := 1; j < len(nums); j++ {
			count := searchAndCount(j, q, nums, sums)
			if maxs[i] < count {
				maxs[i] = count
			}
		}
	}
	return maxs
}

func searchAndCount(index, limit int, nums, sums []int) int {
	if nums[index] > limit {
		return 0
	}
	if nums[index] == limit {
		return 1
	}
	l := index
	r := len(nums)
	mid := (l + r) / 2
	for l < r {
		mid = (l + r) / 2
		sub := sub(index, mid, nums, sums)
		if sub == limit {
			break
		}
		if sub > limit {
			r = mid
			continue
		}
		if sub < limit {
			l = mid + 1
			continue
		}
	}
	if sub(index, mid, nums, sums) > limit {
		return mid - index
	}
	return mid - index + 1
}

func sub(from, to int, nums, sums []int) int {
	return sums[to] - sums[from] + nums[from]
}
