package main

func closedIsland(grid [][]int) int {
	h := len(grid)
	w := len(grid[0])
	queue := make([]int, 0, 100)
	traveled := make([]bool, h*w)
	count := 0
	var x, y int
	for i := 0; i < h*w; i++ {
		x, y = toPoint(i, h, w)
		if traveled[i] {
			continue
		}
		if !isLand(grid, x, y) {
			traveled[i] = true
			continue
		}
		queue = append(queue, i)
		edge := false
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			x, y = toPoint(cur, h, w)
			traveled[cur] = true
			if !isLand(grid, x, y) {
				continue
			}
			edge = edge || isEdge(x, y, h, w)
			// left
			if x-1 >= 0 && !traveled[(x-1)+y*w] {
				queue = append(queue, (x-1)+y*w)
			}
			// right
			if x+1 < w && !traveled[(x+1)+y*w] {
				queue = append(queue, (x+1)+y*w)
			}
			// up
			if y-1 >= 0 && !traveled[x+(y-1)*w] {
				queue = append(queue, x+(y-1)*w)
			}
			// down
			if y+1 < h && !traveled[x+(y+1)*w] {
				queue = append(queue, x+(y+1)*w)
			}
		}
		if !edge {
			count++
		}
	}
	return count
}

func toPoint(index, h, w int) (x, y int) {
	x = index % w
	y = index / w
	return
}

func isLand(grid [][]int, x, y int) bool {
	return grid[y][x] == 0
}

func isEdge(x, y, h, w int) bool {
	return x == 0 || x == w-1 || y == 0 || y == h-1
}
