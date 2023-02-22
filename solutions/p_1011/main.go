package main

import "fmt"

func main() {
	case1()
}

func case1() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	days := 5
	fmt.Println(shipWithinDays(weights, days))
}
