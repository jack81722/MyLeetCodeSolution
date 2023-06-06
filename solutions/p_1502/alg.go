package main

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	sort.Slice(arr, func(i int, j int) bool {
		return arr[i] < arr[j]
	})
	interval := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != interval {
			return false
		}
	}
	return true
}
