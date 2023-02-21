package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	fmt.Println(singleNonDuplicate(nums))
}

func case2() {
	nums := []int{3, 3, 7, 7, 10, 11, 11}
	fmt.Println(singleNonDuplicate(nums))
}

func case3() {
	nums := []int{1, 1, 2, 3, 3}
	fmt.Println(singleNonDuplicate(nums))
}

func case4() {
	nums := []int{1, 1, 2, 2, 3}
	fmt.Println(singleNonDuplicate(nums))
}
