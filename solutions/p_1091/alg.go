package main

func shortestPathBinaryMatrix(grid [][]int) int {
	size := len(grid)
	mp := make([]int, size*size)
	for i := 0; i < len(mp); i++ {
		mp[i] = -1
	}
	if grid[0][0] == 0 {
		mp[0] = 1
	}
	closed := make([]bool, size*size)
	queue := make([]int, 0, size)
	root := toIndex(0, 0, size)
	queue = append(queue, root)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		neighbors := neighbor(grid, cur, size)
		for _, n := range neighbors {
			x, y := toPoint(n, size)
			if closed[n] || grid[y][x] != 0 {
				continue
			}
			if mp[cur]+1 < mp[n] || mp[n] < 0 {
				mp[n] = mp[cur] + 1
			}

			queue = append(queue, n)
			closed[n] = true
		}
	}
	return mp[size*size-1]
}

func toIndex(x, y, size int) int {
	return x + y*size
}

func toPoint(index, size int) (int, int) {
	x := index % size
	y := index / size
	return x, y
}

func neighbor(grid [][]int, index, size int) []int {
	x, y := toPoint(index, size)
	results := make([]int, 0, 8)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			idx := (x + i) + (y+j)*size
			if x+i >= 0 && x+i < size && y+j >= 0 && y+j < size && idx != index && grid[y][x] == 0 {
				results = append(results, idx)
			}
		}
	}
	return results
}
