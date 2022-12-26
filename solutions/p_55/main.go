package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
}

func case2() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}

func case3() {
	nums := []int{0, 2, 3}
	fmt.Println(canJump(nums))
}
