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
	tree := scenarios.NewTree(4, 2, 6, 1, 3)
	fmt.Println(minDiffInBST(tree))
}

func case2() {
	tree := scenarios.NewTree(10, 0, 48, nil, nil, 12)
	fmt.Println(minDiffInBST(tree))
}
