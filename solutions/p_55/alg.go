package main

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	reach := make([]bool, len(nums))
	reach[0] = true
	for i, n := range nums {
		if !reach[i] {
			continue
		}
		if i+n >= len(nums)-1 {
			return true
		}
		for j := 1; j <= n && i+j < len(nums); j++ {
			reach[i+j] = true
		}
	}
	return reach[len(nums)-1]
}
