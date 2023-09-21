package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j int
	n, m := len(nums1), len(nums2)
	merged := make([]int, 0, n+m)
	for i < n || j < m {
		if i >= n {
			merged = append(merged, nums2[j])
			j++
			continue
		}
		if j >= m {
			merged = append(merged, nums1[i])
			i++
			continue
		}
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else if nums1[i] > nums2[j] {
			merged = append(merged, nums2[j])
			j++
		} else {
			merged = append(merged, nums1[i], nums2[j])
			i++
			j++
		}
	}
	mid := (m + n) / 2
	if (m+n)%2 == 1 {
		return float64(merged[mid])
	}
	return float64(merged[mid]+merged[mid-1]) / 2
}
