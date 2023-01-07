package main

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	global, cur, start := 0, 0, 0
	for i := 0; i < n; i++ {
		global += gas[i] - cost[i]
		cur += gas[i] - cost[i]
		if cur < 0 {
			start = i + 1
			cur = 0
		}
	}
	if global < 0 {
		return -1
	}
	return start
}
