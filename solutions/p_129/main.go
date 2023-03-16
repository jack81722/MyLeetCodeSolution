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
	root.Left = scenarios.NewTreeNode(2)
	root.Right = scenarios.NewTreeNode(3)
	fmt.Println(sumNumbers(root))
}

func case2() {
	root := scenarios.NewTreeNode(4)
	root.Left = scenarios.NewTreeNode(9)
	root.Left.Left = scenarios.NewTreeNode(5)
	root.Left.Right = scenarios.NewTreeNode(1)
	root.Right = scenarios.NewTreeNode(0)
	fmt.Println(sumNumbers(root))
}
