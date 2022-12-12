package main

func minCostClimbingStairs2(cost []int) int {
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	i := 2
	for ; i < len(cost); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	return min(dp[i-1], dp[i-2])
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
