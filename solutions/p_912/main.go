package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	nums := []int{5, 2, 3, 1}
	fmt.Println(sortArray(nums))
}

func case2() {
	nums := []int{5, 1, 1, 2, 0, 0}
	fmt.Println(sortArray(nums))
}

func case3() {
	nums := []int{-2, 3, -5}
	fmt.Println(sortArray(nums))
}

func case4() {
	nums := []int{-7, -5, -4, -1, -1, 0, 0, 4, 7, 9}
	fmt.Println(sortArray(nums))
}
