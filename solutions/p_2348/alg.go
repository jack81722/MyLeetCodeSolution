package main

func zeroFilledSubarray(nums []int) int64 {
	var total int64
	var count int64
	for _, n := range nums {
		if n == 0 {
			count++
			total += count
			continue
		}
		if n != 0 && count > 0 {
			count = 0
		}
	}
	return total
}
