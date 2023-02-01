package main

import (
	"sort"
)

func bestTeamScore(scores []int, ages []int) int {
	// sort
	team := make([][]int, len(scores))
	for i := 0; i < len(scores); i++ {
		team[i] = []int{ages[i], scores[i]}
	}
	sort.Slice(team, func(i, j int) bool {
		if team[i][0] == team[j][0] {
			return team[i][1] < team[j][1]
		}
		return team[i][0] < team[j][0]
	})
	//
	dp := make([]int, len(scores))
	dp[0] = team[0][1]
	for i := 1; i < len(scores); i++ {
		dp[i] = team[i][1]
		for j := 0; j < i; j++ {
			if team[j][1] > team[i][1] {
				continue
			}
			dp[i] = max(dp[i], team[i][1]+dp[j])
		}
	}
	return max(dp...)

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
