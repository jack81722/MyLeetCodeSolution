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
	root := scenarios.NewTree(4, 2, 6, 1, 3)
	fmt.Println(getMinimumDifference(root))
}

func case2() {
	root := scenarios.NewTree(1, 0, 48, nil, nil, 12, 49)
	fmt.Println(getMinimumDifference(root))
}
