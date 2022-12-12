package main

func fib(n int) int {
	if n < 2 {
		return n
	}
	sum := 0
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		sum = a + b
		a = b
		b = sum
	}
	return sum
}
