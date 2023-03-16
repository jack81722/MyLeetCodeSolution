package main

import (
	"fmt"
	"leetcode/scenarios"
)

func main() {
	// case1()
	// case2()
	// case3()
	// case4()
	case5()
}

func case1() {
	root := scenarios.NewTree(1, 2, 3, 4, 5, 6)
	fmt.Println(isCompleteTree(root))
}

func case2() {
	root := scenarios.NewTree(1, 2, 3, 4, 5, nil, 7)
	fmt.Println(isCompleteTree(root))
}

func case3() {
	root := scenarios.NewTree(1, 2, 3, 5, nil, 7, 8)
	fmt.Println(isCompleteTree(root))
}

func case4() {
	root := scenarios.NewTree(1, 2, 3, 5, nil, 7)
	fmt.Println(isCompleteTree(root))
}

func case5() {
	root := scenarios.NewTree(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, nil, nil, 15)
	fmt.Println(isCompleteTree(root))
}
