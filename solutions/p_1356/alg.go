package main

import "sort"

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		a, b := bits(arr[i]), bits(arr[j])
		if a == b {
			return arr[i] < arr[j]
		}
		return a < b
	})
	return arr
}

func bits(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
