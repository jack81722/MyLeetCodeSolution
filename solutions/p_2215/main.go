package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	fmt.Println(findDifference(nums1, nums2))
}

func case2() {
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{1, 1, 2, 2}
	fmt.Println(findDifference(nums1, nums2))
}
