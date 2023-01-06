package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	costs := []int{1, 3, 2, 4, 1}
	coins := 7
	fmt.Println(maxIceCream(costs, coins))
}

func case2() {
	costs := []int{10, 6, 8, 7, 7, 8}
	coins := 5
	fmt.Println(maxIceCream(costs, coins))
}

func case3() {
	costs := []int{1, 6, 3, 1, 2, 5}
	coins := 20
	fmt.Println(maxIceCream(costs, coins))
}
