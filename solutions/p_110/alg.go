package main

import "leetcode/scenarios"

func isBalanced(root *scenarios.TreeNode) bool {
	_, b := getHeightAndBalanced(root)
	return b
}

func getHeightAndBalanced(root *scenarios.TreeNode) (h int, b bool) {
	if root == nil {
		return 0, true
	}
	lh, lb := getHeightAndBalanced(root.Left)
	rh, rb := getHeightAndBalanced(root.Right)
	if !lb || !rb {
		return max(lh, rh) + 1, false
	}
	return max(lh, rh) + 1, max(lh, rh)-min(lh, rh) < 2
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
