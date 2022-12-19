package main

func countPrimes(n int) int {
	visited := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if visited[i] {
			continue
		}
		count++
		for j := 2; j*i < n; j++ {
			visited[j*i] = true
		}
	}
	return count
}
