package main

// my original solution
func singleNonDuplicate2(nums []int) int {
	l, r := 0, len(nums)-1
	var count, half int
	for {
		count = r - l + 1
		if count == 1 {
			return nums[l]
		}
		half = count/2 + l
		if nums[half-1] != nums[half] && nums[half+1] != nums[half] {
			return nums[half]
		} else if nums[half-1] == nums[half] {
			if (half-1)%2 == 1 {
				r = half - 2
				continue
			}
			l = half + 1
			continue
		}
		if half%2 == 1 {
			r = half - 1
			continue
		}
		l = half + 2
	}
}
