package main

func jump(nums []int) int {
	count := make([]int, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		for j := 1; j <= nums[i] && i+j < len(nums); j++ {
			if count[i+j] == 0 {
				count[i+j] = count[i] + 1
				continue
			}
			count[i+j] = min(count[i]+1, count[i+j])
		}
	}
	return count[len(nums)-1]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
