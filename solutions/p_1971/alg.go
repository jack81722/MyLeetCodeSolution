package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	// map
	mp := make([][]int, n)
	for _, arr := range edges {
		mp[arr[0]] = append(mp[arr[0]], arr[1])
		mp[arr[1]] = append(mp[arr[1]], arr[0])
	}
	return step(source, destination, mp, make([]bool, n))
}

func step(index, dest int, mp [][]int, closed []bool) bool {
	if index == dest {
		return true
	}
	closed[index] = true

	nexts := mp[index]
	for _, next := range nexts {
		if closed[next] {
			continue
		}
		if step(next, dest, mp, closed) {
			return true
		}
	}
	return false
}
