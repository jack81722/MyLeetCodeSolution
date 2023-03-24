package main

func minReorder(n int, connections [][]int) int {
	graph := make([][]int, n)
	for _, conn := range connections {
		graph[conn[0]] = append(graph[conn[0]], -conn[1])
		graph[conn[1]] = append(graph[conn[1]], conn[0])
	}
	count := 0
	closed := make([]bool, n)
	closed[0] = true
	queue := make([]int, 0, n)
	queue = append(queue, 0)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, next := range graph[cur] {
			neg := next < 0
			abs := next
			if neg {
				abs = -abs
			}
			if closed[abs] {
				continue
			}
			if neg {
				count++
			}
			closed[abs] = true
			queue = append(queue, abs)
		}
	}
	return count
}
