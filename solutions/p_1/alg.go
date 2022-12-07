package main

func twoSum(nums []int, target int) []int {
	for i, n := range nums {
		for j := i + 1; j < len(nums); j++ {
			if target == n+nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}
