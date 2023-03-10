package main

import (
	"fmt"
	"leetcode/scenarios"
)

func main() {
	case1()
}

func case1() {
	list := scenarios.NewList(1, 2, 3)
	sol := Constructor(list)
	fmt.Println(sol.GetRandom())
	fmt.Println(sol.GetRandom())
	fmt.Println(sol.GetRandom())
	fmt.Println(sol.GetRandom())
	fmt.Println(sol.GetRandom())
}
