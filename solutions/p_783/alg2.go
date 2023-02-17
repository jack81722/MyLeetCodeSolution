package main

import (
	"leetcode/scenarios"
)

func minDiffInBST2(root *scenarios.TreeNode) int {
	result := make([]int, 0, 1000)
	traversal(root, &result)
	m := result[1] - result[0]
	for i := 1; i < len(result)-1; i++ {
		if result[i+1]-result[i] < m {
			m = result[i+1] - result[i]
		}
	}
	return m
}

func traversal(root *scenarios.TreeNode, result *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		traversal(root.Left, result)
	}
	*result = append(*result, root.Val)
	if root.Right != nil {
		traversal(root.Right, result)
	}
}
