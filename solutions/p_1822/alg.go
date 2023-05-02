package main

func arraySign(nums []int) int {
	result := nums[0] % 2
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0
		}
		if nums[i] < 0 {
			result = -result
		}
	}
	return result
}
