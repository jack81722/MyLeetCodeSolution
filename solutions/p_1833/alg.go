package main

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i] < costs[j]
	})
	total := 0
	count := 0
	for i := 0; i < len(costs) && total+costs[i] <= coins; i++ {
		count++
		total += costs[i]
	}
	return count
}
