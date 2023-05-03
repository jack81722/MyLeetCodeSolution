package main

func findDifference2(nums1 []int, nums2 []int) [][]int {
	mp1 := make(map[int]bool, len(nums1))
	mp2 := make(map[int]bool, len(nums2))
	for i := 0; i < max(len(nums1), len(nums2)); i++ {
		if i < len(nums1) {
			mp1[nums1[i]] = true
		}
		if i < len(nums2) {
			mp2[nums2[i]] = true
		}
	}
	ans1 := make([]int, 0, len(nums1))
	ans2 := make([]int, 0, len(nums2))
	for k := range mp1 {
		if !mp2[k] {
			ans1 = append(ans1, k)
		}
	}
	for k := range mp2 {
		if !mp1[k] {
			ans2 = append(ans2, k)
		}
	}
	return [][]int{ans1, ans2}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
