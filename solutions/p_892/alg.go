package main

func surfaceArea(grid [][]int) int {
	n := len(grid)
	area := 0
	for y, _ := range grid {
		for x, _ := range grid[y] {
			if grid[y][x] == 0 {
				continue
			}
			area += 4*grid[y][x] + 2
			// left
			if inBound(x-1, y, n) {
				area -= min(grid[y][x], grid[y][x-1])
			}
			// right
			if inBound(x+1, y, n) {
				area -= min(grid[y][x], grid[y][x+1])
			}
			// up
			if inBound(x, y-1, n) {
				area -= min(grid[y][x], grid[y-1][x])
			}
			// down
			if inBound(x, y+1, n) {
				area -= min(grid[y][x], grid[y+1][x])
			}
		}
	}
	return area
}

func inBound(x, y, n int) bool {
	return x >= 0 && x < n && y >= 0 && y < n
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
