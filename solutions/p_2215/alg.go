package main

/*
 * Reference: https://leetcode.com/problems/find-the-difference-of-two-arrays/solutions/2825038/go-solution-using-one-hashmap-21ms-faster-than-100/
 */

func findDifference(nums1 []int, nums2 []int) [][]int {
	res := make([][]int, 2)
	res[0] = make([]int, 0, len(nums1))
	res[1] = make([]int, 0, len(nums1))
	m := make(map[int]int, len(nums1)+len(nums2))

	for _, v := range nums1 {
		m[v] = -1
	}
	for _, v := range nums2 {
		m[v] += 2
	}

	for k, v := range m {
		if v == -1 {
			res[0] = append(res[0], k)
		} else if v%2 == 0 {
			res[1] = append(res[1], k)
		}
	}

	return res
}
