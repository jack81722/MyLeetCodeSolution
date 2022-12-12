package main

func climbStairs2(n int) int {
	if n < 3 {
		return n
	}
	sum := 0
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		sum = a + b
		a = b
		b = sum
	}
	return sum
}
