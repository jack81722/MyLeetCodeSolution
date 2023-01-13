package main

func longestPath(parent []int, s string) int {
	graph := make([][]int, len(parent))
	for n, p := range parent {
		if p < 0 {
			continue
		}
		graph[n] = append(graph[n], p)
		graph[p] = append(graph[p], n)
	}
	max := 1
	dp := make([]map[int]int, len(parent))
	for i := range dp {
		dp[i] = map[int]int{}
	}
	for i := range parent {
		m := dfs(graph, parent, s, i, i, 1, dp)
		if max < m {
			max = m
		}
	}
	return max
}

func dfs(graph [][]int, parents []int, labels string, index, from, max int, dp []map[int]int) int {
	m := max
	for _, n := range graph[index] {
		if n == from {
			continue
		}
		if labels[n] == labels[index] {
			continue
		}

		mm, ok := dp[index][n]
		if !ok {
			mm = dfs(graph, parents, labels, n, index, 1, dp)
			dp[index][n] = mm
		}
		if mm+max > m {
			m = mm + max
		}
	}
	return m
}
