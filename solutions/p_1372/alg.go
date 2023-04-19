package main

import (
	"leetcode/scenarios"
)

func longestZigZag(root *scenarios.TreeNode) int {
	_, _, m := dfs(root)
	return m
}

func dfs(root *scenarios.TreeNode) (int, int, int) {
	if root == nil {
		return 0, 0, 0
	}
	var l, r, lm, rm int
	_, l, lm = dfs(root.Left)
	r, _, rm = dfs(root.Right)
	return l + 1, r + 1, max(max(l, r), max(lm, rm))
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
