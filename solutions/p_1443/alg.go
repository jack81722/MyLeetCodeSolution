package main

func minTime(n int, edges [][]int, hasApple []bool) int {
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	from := 0
	return dfs(graph, hasApple, 0, from)
}

func dfs(graph [][]int, hasApple []bool, index, from int) int {
	dist := 0
	for _, next := range graph[index] {
		if next == from {
			continue
		}
		dist += dfs(graph, hasApple, next, index)
	}
	if (dist > 0 || hasApple[index]) && index > 0 {
		dist += 2
	}
	return dist
}
