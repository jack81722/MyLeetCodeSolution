package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
}

func case2() {
	prices := []int{2, 1}
	fmt.Println(maxProfit(prices))
}

func case3() {
	prices := []int{0}
	fmt.Println(maxProfit(prices))
}

func case4() {
	prices := []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}
	fmt.Println(maxProfit(prices))
}
