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
	tree := scenarios.NewTree(3, 9, 20, nil, nil, 15, 7)
	fmt.Println(maxDepth(tree))
}

func case2() {
	tree := scenarios.NewTree(1, nil, 2)
	fmt.Println(maxDepth(tree))
}
