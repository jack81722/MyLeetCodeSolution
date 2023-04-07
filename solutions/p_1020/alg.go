package main

func numEnclaves(grid [][]int) int {
	h := len(grid)
	w := len(grid[0])
	queue := make([]int, 0, 100)
	traveled := make([]bool, w*h)
	var x, y, totalCount int
	for i := 0; i < h*w; i++ {
		if traveled[i] {
			continue
		}
		traveled[i] = true
		x, y = toPoint(i, w, h)
		if grid[y][x] == 0 {
			traveled[i] = true
			continue
		}

		queue = append(queue, i)
		isEdge := false
		count := 0
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			count++
			x, y = toPoint(cur, w, h)
			if !isEdge {
				isEdge = isEdge || (x == 0 || x == w-1 || y == 0 || y == h-1)
			}
			if x-1 >= 0 && grid[y][x-1] == 1 && !traveled[cur-1] {
				queue = append(queue, cur-1)
				traveled[cur-1] = true
			}
			if x+1 < w && grid[y][x+1] == 1 && !traveled[cur+1] {
				queue = append(queue, cur+1)
				traveled[cur+1] = true
			}
			if y-1 >= 0 && grid[y-1][x] == 1 && !traveled[cur-w] {
				queue = append(queue, cur-w)
				traveled[cur-w] = true
			}
			if y+1 < h && grid[y+1][x] == 1 && !traveled[cur+w] {
				queue = append(queue, cur+w)
				traveled[cur+w] = true
			}
		}
		x, y = toPoint(i, w, h)
		if !isEdge {
			totalCount += count
		}
	}
	return totalCount
}

func toPoint(index, w, h int) (x, y int) {
	x = index % w
	y = index / w
	return
}
