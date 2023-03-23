package main

func makeConnected(n int, connections [][]int) int {
	graph := make([][]int, n)
	for _, conn := range connections {
		graph[conn[0]] = append(graph[conn[0]], conn[1])
		graph[conn[1]] = append(graph[conn[1]], conn[0])
	}
	count := len(connections)
	need := -1
	closed := make([]bool, n)
	queue := make([]int, 0, len(connections))
	for i := 0; i < n; i++ {
		if closed[i] {
			continue
		}
		need++
		closed[i] = true
		queue = append(queue, i)
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, next := range graph[cur] {
				if closed[next] {
					continue
				}
				closed[next] = true
				queue = append(queue, next)
				count--
			}
		}
	}
	if need > count {
		return -1
	}
	return need
}
