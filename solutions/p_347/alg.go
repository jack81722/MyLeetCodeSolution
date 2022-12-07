package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	mapped := map[int]int{}
	for _, n := range nums {
		mapped[n] += 1
	}
	keys := make([]int, len(mapped))
	i := 0
	for key, _ := range mapped {
		keys[i] = key
		i++
	}
	sort.Slice(keys, func(i, j int) bool {
		return mapped[keys[i]] > mapped[keys[j]]
	})
	result := make([]int, k)
	for dk := 0; dk < k; dk++ {
		result[dk] = keys[dk]
	}
	return result
}
