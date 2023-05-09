package main

func spiralOrder(matrix [][]int) []int {
	vectors := []Vector{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	h := len(matrix)
	w := len(matrix[0])
	closed := make([]bool, h*w)
	result := make([]int, 0, h*w)
	p := Vector{x: 0, y: 0}
	var vi, count int
	v := vectors[vi]
	for count < h*w {
		result = append(result, matrix[p.y][p.x])
		closed[p.toIdx(w)] = true
		count++
		next := Vector{x: p.x + v.x, y: p.y + v.y}
		if next.x < 0 || next.x >= w || next.y < 0 || next.y >= h || closed[next.toIdx(w)] {
			vi = (vi + 1) % len(vectors)
			v = vectors[vi]
		}
		p.x += v.x
		p.y += v.y
	}
	return result
}

type Vector struct {
	x, y int
}

func (v Vector) toIdx(w int) int {
	return v.y*w + v.x
}
