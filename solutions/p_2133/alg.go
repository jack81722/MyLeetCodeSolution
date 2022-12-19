package main

func checkValid(matrix [][]int) bool {
	l := len(matrix)
	rowMap := newMap(l)
	colMap := newMap(l)
	for y, _ := range matrix {
		for x, _ := range matrix {
			n := matrix[y][x] - 1
			rowMap[y][n] = true
			colMap[x][n] = true
		}
	}
	for _, r := range rowMap {
		if !allTrue(r) {
			return false
		}
	}
	for _, c := range colMap {
		if !allTrue(c) {
			return false
		}
	}
	return true
}

func newMap(n int) [][]bool {
	m := make([][]bool, n)
	for i, _ := range m {
		m[i] = make([]bool, n)
	}
	return m
}

func allTrue(list []bool) bool {
	for _, b := range list {
		if !b {
			return false
		}
	}
	return true
}
