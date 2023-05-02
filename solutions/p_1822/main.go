package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	nums := []int{-1, -2, -3, -4, 3, 2, 1}
	fmt.Println(arraySign(nums))
}

func case2() {
	nums := []int{1, 5, 0, 2, -3}
	fmt.Println(arraySign(nums))
}

func case3() {
	nums := []int{-1, 1, -1, 1, -1}
	fmt.Println(arraySign(nums))
}

func case4() {
	nums := []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}
	fmt.Println(arraySign(nums))
}
