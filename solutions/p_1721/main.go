package main

import (
	"fmt"
	"leetcode/scenarios"
)

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	head := scenarios.NewList(1, 2, 3, 4, 5)
	k := 2
	fmt.Println(swapNodes(head, k).String())
}

func case2() {
	head := scenarios.NewList(7, 9, 6, 6, 7, 8, 3, 0, 9, 5)
	k := 5
	fmt.Println(swapNodes(head, k).String())
}

func case3() {
	head := scenarios.NewList(1, 2)
	k := 1
	fmt.Println(swapNodes(head, k).String())
}

func case4() {
	head := scenarios.NewList(1, 2)
	k := 2
	fmt.Println(swapNodes(head, k).String())
}
