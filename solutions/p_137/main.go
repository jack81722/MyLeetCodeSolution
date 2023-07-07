package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	nums := []int{2, 2, 3, 2}
	fmt.Println(singleNumber(nums))
}

func case2() {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber(nums))
}
