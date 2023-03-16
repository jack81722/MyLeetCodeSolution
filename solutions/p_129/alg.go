package main

import (
	"leetcode/scenarios"
)

func sumNumbers(root *scenarios.TreeNode) int {
	if root == nil {
		return 0
	}
	return sumTravel(0, 0, root)
}

func sumTravel(sum, num int, root *scenarios.TreeNode) int {
	num = num*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum + num
	}
	if root.Left != nil {
		sum = sumTravel(sum, num, root.Left)
	}
	if root.Right != nil {
		sum = sumTravel(sum, num, root.Right)
	}
	return sum
}
