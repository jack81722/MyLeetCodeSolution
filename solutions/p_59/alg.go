package main

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	i := 1
	var curX, curY, nextX, nextY, vi int
	vectors := [][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	sqrN := n * n
	for i <= sqrN {
		matrix[curY][curX] = i
		i++
		nextX, nextY = curX+vectors[vi][0], curY+vectors[vi][1]
		if nextX < 0 || nextX >= n || nextY < 0 || nextY >= n || matrix[nextY][nextX] != 0 {
			vi = (vi + 1) % len(vectors)
			nextX, nextY = curX+vectors[vi][0], curY+vectors[vi][1]
		}
		curX, curY = nextX, nextY
	}
	return matrix
}
