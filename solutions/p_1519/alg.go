package main

func countSubTrees(n int, edges [][]int, labels string) []int {
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	result := make([]int, n)
	count := make([]int, 26)
	var dfs func(index, from int)
	dfs = func(index, from int) {
		atoi := byte(labels[index] - 'a')
		recent := count[atoi]
		count[atoi]++
		for _, n := range graph[index] {
			if n == from {
				continue
			}
			dfs(n, index)
		}
		result[index] = count[atoi] - recent
	}
	dfs(0, 0)
	return result
}
