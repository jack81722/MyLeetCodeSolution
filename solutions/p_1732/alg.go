package main

func largestAltitude(gain []int) int {
	cur := 0
	max := 0
	for _, g := range gain {
		cur = cur + g
		if cur > max {
			max = cur
		}
	}
	return max
}
