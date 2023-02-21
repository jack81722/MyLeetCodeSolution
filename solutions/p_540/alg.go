package main

func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	var count, half int
	for l < r {
		count = r - l
		half = count/2 + l
		if nums[half] == nums[half^1] {
			l = half + 1
		} else {
			r = half
		}
	}
	return nums[l]
}
