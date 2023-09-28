package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	nums := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(nums))
}

func case2() {
	nums := []int{0, 2, 1}
	fmt.Println(sortArrayByParity(nums))
}
