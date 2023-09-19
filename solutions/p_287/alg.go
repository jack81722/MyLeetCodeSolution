package main

func findDuplicate(nums []int) int {
	mp := map[int]bool{}
	for _, n := range nums {
		if mp[n] {
			return n
		}
		mp[n] = true
	}
	return -1
}
