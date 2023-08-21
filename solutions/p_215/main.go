package main

import "log"

func main() {
	case1()
	case2()
}

func case1() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	log.Println(findKthLargest(nums, k))
}

func case2() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	log.Println(findKthLargest(nums, k))
}
