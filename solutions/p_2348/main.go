package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	nums := []int{1, 3, 0, 0, 2, 0, 0, 4}
	fmt.Println(zeroFilledSubarray(nums))
}

func case2() {
	nums := []int{0, 0, 0, 2, 0, 0}
	fmt.Println(zeroFilledSubarray(nums))
}

func case3() {
	nums := []int{2, 10, 2019}
	fmt.Println(zeroFilledSubarray(nums))
}
