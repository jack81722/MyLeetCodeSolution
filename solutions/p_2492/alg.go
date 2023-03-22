package main

import "math"

func minScore(n int, roads [][]int) int {
	graph := make([][][2]int, n+1)
	for _, road := range roads {
		graph[road[0]] = append(graph[road[0]], [2]int{road[1], road[2]})
		graph[road[1]] = append(graph[road[1]], [2]int{road[0], road[2]})
	}

	minDist := math.MaxInt
	closed := map[int]bool{1: true}
	queue := make([]int, 0, n)
	queue = append(queue, 1)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, dest := range graph[cur] {
			minDist = min(minDist, dest[1])
			if closed[dest[0]] {
				continue
			}
			queue = append(queue, dest[0])
			closed[dest[0]] = true
		}
	}
	return minDist
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
