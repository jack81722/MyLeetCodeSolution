package main

func maximumDifference(nums []int) int {
	minN := nums[0]
	maxAns := -1
	for _, n := range nums {
		if minN >= n {
			minN = n
		} else if maxAns < n-minN {
			maxAns = n - minN
		}
	}
	return maxAns
}
