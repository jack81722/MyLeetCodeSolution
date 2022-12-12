package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost))
}

func case2() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}
