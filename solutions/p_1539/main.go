package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	arr := []int{2, 3, 4, 7, 11}
	k := 5
	fmt.Println(findKthPositive(arr, k))
}

func case2() {
	arr := []int{1, 2, 3, 4}
	k := 2
	fmt.Println(findKthPositive(arr, k))
}
