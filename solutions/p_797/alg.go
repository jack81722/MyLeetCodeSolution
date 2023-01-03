package main

func allPathsSourceTarget(graph [][]int) [][]int {
	passed := make([]int, 0, len(graph))
	result := make([][]int, 0, 100)
	closed := make([]bool, len(graph))
	_ = step(0, passed, graph, &result, closed)
	return result
}

func step(index int, passed []int, graph [][]int, result *[][]int, closed []bool) []int {
	n := len(graph)
	passed = append(passed, index)
	closed[index] = true
	for _, next := range graph[index] {
		if next == n-1 {
			r := make([]int, len(passed)+1)
			copy(r, passed)
			r[len(r)-1] = next
			*result = append(*result, r)
			continue
		}
		if !closed[next] {
			passed = step(next, passed, graph, result, closed)
		}
	}
	passed = passed[:len(passed)-1]
	closed[index] = false
	return passed
}
