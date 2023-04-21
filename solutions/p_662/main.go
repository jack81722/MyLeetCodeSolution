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
	root.Left = scenarios.NewTreeNode(3)
	root.Right = scenarios.NewTreeNode(2)
	root.Left.Left = scenarios.NewTreeNode(5)
	root.Left.Right = scenarios.NewTreeNode(3)
	root.Right.Right = scenarios.NewTreeNode(9)
	fmt.Println(widthOfBinaryTree(root))
}

func case2() {
	root := scenarios.NewTreeNode(1)
	root.Left = scenarios.NewTreeNode(3)
	root.Right = scenarios.NewTreeNode(2)
	root.Right.Right = scenarios.NewTreeNode(9)
	root.Right.Right.Left = scenarios.NewTreeNode(7)
	root.Left.Left = scenarios.NewTreeNode(5)
	root.Left.Left.Left = scenarios.NewTreeNode(6)
	fmt.Println(widthOfBinaryTree(root))
}
