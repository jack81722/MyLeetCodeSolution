package main

import "sort"

// brute-force
func canCompleteCircuit2(gas []int, cost []int) int {
	n := len(gas)
	candidate := make([]int, 0, 100)
	for i, c := range cost {
		gas[i] -= c
		if gas[i] >= 0 {
			candidate = append(candidate, i)
		}
	}
	sort.Slice(candidate, func(i, j int) bool {
		return candidate[i] > candidate[j]
	})
	for len(candidate) > 0 {
		cur := candidate[0]
		candidate = candidate[1:]
		remain := gas[cur]
		for i := next(cur, n); i != cur; i = next(i, n) {
			remain += gas[i]
			if remain < 0 {
				break
			}
		}
		if remain >= 0 {
			return cur
		}
	}
	return -1
}

func next(i, n int) int {
	return (i + 1) % n
}
