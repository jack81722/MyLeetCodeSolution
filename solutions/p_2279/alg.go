package main

import (
	"sort"
)

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	for i, c := range capacity {
		capacity[i] = c - rocks[i]
	}
	sort.Slice(capacity, func(i, j int) bool {
		return capacity[i] < capacity[j]
	})
	i := 0
	for i < len(capacity) && additionalRocks-capacity[i] >= 0 {
		additionalRocks -= capacity[i]
		i++
	}
	return i
}
