package main

import (
	"fmt"
	"leetcode/scenarios"
)

func main() {
	case1()
}

func case1() {
	root := scenarios.NewTree(4, 2, 7, 1, 3, 6, 9)
	fmt.Println(invertTree(root))
}
