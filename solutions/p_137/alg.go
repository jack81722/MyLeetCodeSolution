package main

func singleNumber(nums []int) int {
	mp := make(map[int]int, (len(nums)-1)/3)
	for _, n := range nums {
		mp[n]++
	}
	for n, c := range mp {
		if c == 1 {
			return n
		}
	}
	return 0
}
