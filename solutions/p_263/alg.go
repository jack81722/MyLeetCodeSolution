package main

func isUgly(n int) bool {
	n = do(n, 2)
	n = do(n, 3)
	n = do(n, 5)
	return n == 1
}

func do(n, div int) int {
	if n == 0 {
		return n
	}
	for n%div == 0 {
		n /= div
	}
	return n
}
