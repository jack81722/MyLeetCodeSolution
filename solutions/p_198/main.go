package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}

func case2() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}

func case3() {
	nums := []int{2, 7, 9, 3, 1, 5}
	fmt.Println(rob(nums))
}
