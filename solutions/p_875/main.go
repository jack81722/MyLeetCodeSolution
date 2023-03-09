package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	piles := []int{3, 6, 7, 11}
	h := 8
	fmt.Println(minEatingSpeed(piles, h))
}

func case2() {
	piles := []int{30, 11, 23, 4, 20}
	h := 5
	fmt.Println(minEatingSpeed(piles, h))
}

func case3() {
	piles := []int{30, 11, 23, 4, 20}
	h := 6
	fmt.Println(minEatingSpeed(piles, h))
}
