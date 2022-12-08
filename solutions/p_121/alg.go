package main

func maxProfit(prices []int) int {
	minB := prices[0]
	maxS := 0
	for _, p := range prices {
		if minB >= p {
			minB = p
		} else if maxS < p-minB {
			maxS = p - minB
		}
	}
	return maxS
}
