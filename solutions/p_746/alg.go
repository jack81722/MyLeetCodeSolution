package main

func minCostClimbingStairs(cost []int) int {
	dq := map[int]int{}
	return min(step(0, cost, dq), step(1, cost, dq))
}

func step(index int, cost []int, dq map[int]int) int {
	if index >= len(cost) {
		return 0
	}
	_, ok := dq[index]
	if ok {
		return dq[index]
	}
	total := min(step(index+1, cost, dq), step(index+2, cost, dq)) + cost[index]
	dq[index] = total
	return total
}
