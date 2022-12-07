package main

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		c := (l + r) / 2
		if nums[c] > target {
			r = c
		}
		if nums[c] < target {
			l = c + 1
		}
		if nums[c] == target {
			return c
		}
	}
	return l
}
