package main

func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	sums, counts := make([]int, n), make([]int, n)

	dfsOrigin(-1, 0, graph, sums, counts)
	dfsOther(-1, 0, n, graph, sums, counts)
	return sums
}

func dfsOrigin(parent, index int, graph [][]int, sums, counts []int) {
	for _, next := range graph[index] {
		if next == parent {
			continue
		}
		dfsOrigin(index, next, graph, sums, counts)
		sums[index] += sums[next] + counts[next]
		counts[index] += counts[next]
	}
	counts[index]++
}

func dfsOther(parent, index, n int, graph [][]int, sums, counts []int) {
	for _, next := range graph[index] {
		if next == parent {
			continue
		}
		sums[next] = sums[index] - counts[next]*2 + n
		dfsOther(index, next, n, graph, sums, counts)
	}
}
