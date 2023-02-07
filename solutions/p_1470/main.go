package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	fmt.Println(shuffle(nums, n))
}

func case2() {
	nums := []int{1, 2, 3, 4, 4, 3, 2, 1}
	n := 4
	fmt.Println(shuffle(nums, n))
}

func case3() {
	nums := []int{1, 1, 2, 2}
	n := 2
	fmt.Println(shuffle(nums, n))
}
