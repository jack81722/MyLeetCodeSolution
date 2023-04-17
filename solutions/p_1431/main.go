package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	candies := []int{2, 3, 5, 1, 3}
	extra := 3
	fmt.Println(kidsWithCandies(candies, extra))
}

func case2() {
	candies := []int{4, 2, 1, 1, 2}
	extra := 1
	fmt.Println(kidsWithCandies(candies, extra))
}
