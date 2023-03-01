package main

func sortArray(nums []int) []int {
	quick(nums, 0, len(nums)-1)
	return nums
}

func quick(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j := l+1, r
	for i <= j {
		if nums[i] <= nums[l] {
			i++
		} else if nums[j] >= nums[l] {
			j--
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			j--
			i++
		}
	}
	nums[j], nums[l] = nums[l], nums[j]
	quick(nums, l, j-1)
	quick(nums, j+1, r)
}
