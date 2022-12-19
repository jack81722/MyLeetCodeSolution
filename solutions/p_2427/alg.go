package main

func commonFactors(a int, b int) int {
	m := min(a, b)
	count := 0
	for i := 1; i <= m; i++ {
		if a%i == 0 && b%i == 0 {
			count++
		}
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
