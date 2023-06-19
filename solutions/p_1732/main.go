package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	gain := []int{-5, 1, 5, 0, -7}
	fmt.Println(largestAltitude(gain))
}

func case2() {
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	fmt.Println(largestAltitude(gain))
}

func case3() {
	gain := []int{52, -91, 72}
	fmt.Println(largestAltitude(gain))
}
