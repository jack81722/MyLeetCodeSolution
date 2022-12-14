package main

func rob(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	if len(nums) > 2 {
		dp[2] = dp[0] + nums[2]
	}
	for i := 3; i < len(nums); i++ {
		dp[i] = max(dp[i-2], dp[i-3]) + nums[i]
	}
	return max(dp[len(nums)-1], dp[len(nums)-2])
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
