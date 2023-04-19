package main

import (
	"fmt"
	"leetcode/scenarios"
)

func main() {
	case1()
	case2()
}

func case1() {
	root := scenarios.NewTreeNode(1)
	root.Right = scenarios.NewTreeNode(2)
	root.Right.Left = scenarios.NewTreeNode(3)
	root.Right.Right = scenarios.NewTreeNode(4)
	root.Right.Right.Left = scenarios.NewTreeNode(5)
	root.Right.Right.Right = scenarios.NewTreeNode(6)
	root.Right.Right.Left.Right = scenarios.NewTreeNode(7)
	root.Right.Right.Left.Right.Right = scenarios.NewTreeNode(8)

	fmt.Println(longestZigZag(root))
}

func case2() {
	root := scenarios.NewTreeNode(1)
	root.Left = scenarios.NewTreeNode(2)
	root.Right = scenarios.NewTreeNode(3)
	root.Left.Right = scenarios.NewTreeNode(4)
	root.Left.Right.Left = scenarios.NewTreeNode(5)
	root.Left.Right.Right = scenarios.NewTreeNode(6)
	root.Left.Right.Left.Right = scenarios.NewTreeNode(7)

	fmt.Println(longestZigZag(root))
}
