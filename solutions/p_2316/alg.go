package main

func countPairs(n int, edges [][]int) int64 {
	graph := make([][]int, n)
	for _, conn := range edges {
		graph[conn[0]] = append(graph[conn[0]], conn[1])
		graph[conn[1]] = append(graph[conn[1]], conn[0])
	}
	result := int64(0)
	totalCount := int64(0)
	closed := make([]bool, n)
	queue := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if closed[i] {
			continue
		}
		count := int64(0)
		queue = append(queue, i)
		closed[i] = true
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			count++
			for _, next := range graph[cur] {
				if closed[next] {
					continue
				}
				closed[next] = true
				queue = append(queue, next)
			}
		}
		result += totalCount * count
		totalCount += count
	}
	return result
}
