package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func case2() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func case3() {
	nums1 := []int{2}
	nums2 := []int{}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
