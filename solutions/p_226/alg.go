package main

import "leetcode/scenarios"

func invertTree(root *scenarios.TreeNode) *scenarios.TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
