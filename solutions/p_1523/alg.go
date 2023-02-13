package main

func countOdds(low int, high int) int {
	sub := high - low + 1
	count := sub/2 + (sub%2)*(low%2)
	return count
}
