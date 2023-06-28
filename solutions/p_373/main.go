package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}

func case2() {
	nums1 := []int{1, 1, 2}
	nums2 := []int{1, 2, 3}
	k := 2
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}

func case3() {
	nums1 := []int{1, 2}
	nums2 := []int{3}
	k := 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}
