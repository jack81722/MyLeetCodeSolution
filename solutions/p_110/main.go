package main

import (
	"fmt"
	"leetcode/scenarios"
)

func main() {
	case1()
}

func case1() {
	tree := scenarios.NewTree(3, 9, 20, nil, nil, 15, 7)
	fmt.Println(isBalanced(tree))
}
