package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}

func case2() {
	nums := []int{2, 3, 0, 1, 4}
	fmt.Println(jump(nums))
}

func case3() {
	nums := []int{1, 2, 0, 1}
	fmt.Println(jump(nums))
}
