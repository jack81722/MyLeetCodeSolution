package main

func maxDistance(grid [][]int) int {
	size := len(grid)
	queue := make([]int, 0, size*size)
	distMp := make([]int, size*size)
	closeMp := make([]bool, size*size)
	for j, row := range grid {
		for i, isLand := range row {
			if isLand == 1 {
				queue = append(queue, j*size+i)
				closeMp[j*size+i] = true
			}
		}
	}
	if len(queue) == 0 || len(queue) == size*size {
		return -1
	}
	var cur int
	for len(queue) > 0 {
		cur = queue[0]
		queue = queue[1:]
		for _, neighbor := range neighbor(cur, size) {
			if distMp[neighbor] > 0 {
				distMp[neighbor] = min(distMp[neighbor], distMp[cur]+1)
			} else if !closeMp[neighbor] {
				distMp[neighbor] = distMp[cur] + 1
			}
			if !closeMp[neighbor] {
				queue = append(queue, neighbor)
				closeMp[neighbor] = true
			}
		}
	}
	return max(distMp...)
}

func neighbor(index int, size int) []int {
	neighbors := make([]int, 0, 4)
	x := index % size
	y := index / size
	if x-1 >= 0 {
		neighbors = append(neighbors, index-1)
	}
	if x+1 < size {
		neighbors = append(neighbors, index+1)
	}
	if y-1 >= 0 {
		neighbors = append(neighbors, index-size)
	}
	if y+1 < size {
		neighbors = append(neighbors, index+size)
	}
	return neighbors
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(n ...int) int {
	m := n[0]
	for i := 1; i < len(n); i++ {
		if m < n[i] {
			m = n[i]
		}
	}
	return m
}
