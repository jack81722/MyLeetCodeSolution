package main

import (
	"leetcode/scenarios"
	"math"
)

func getMinimumDifference(root *scenarios.TreeNode) int {
	minDiff, prevVal := math.MaxInt, -1
	dfs(root, &minDiff, &prevVal)
	return minDiff
}

func dfs(root *scenarios.TreeNode, minDiff, prevVal *int) {
	if root == nil {
		return
	}

	dfs(root.Left, minDiff, prevVal)

	diff := root.Val - *prevVal
	if *prevVal >= 0 && diff < *minDiff {
		*minDiff = diff
	}
	*prevVal = root.Val

	dfs(root.Right, minDiff, prevVal)
}
