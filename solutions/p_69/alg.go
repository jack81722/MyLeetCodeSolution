package main

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	left := 0
	right := x
	for left < right {
		mid := (left + right) / 2
		if x/mid >= mid {
			left = mid + 1
			continue
		}
		right = mid
	}
	return right - 1
}
