package main

import (
	"math"
)

func maxProfit(prices []int) int {
	sold := 0
	rest := 0
	hold := math.MinInt
	for _, p := range prices {
		prev := sold
		sold = hold + p
		hold = max(hold, rest-p)
		rest = max(rest, prev)
	}
	return max(sold, rest)
}

func max(n ...int) int {
	m := n[0]
	for i := 1; i < len(n); i++ {
		if m < n[i] {
			m = n[i]
		}
	}
	return m
}
