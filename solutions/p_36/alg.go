package main

func isValidSudoku(board [][]byte) bool {
	colMap := newMap()
	rowMap := newMap()
	boxMap := newMap()
	for y, _ := range board {
		for x, _ := range board[y] {
			if board[y][x] == Empty {
				continue
			}
			n := int(board[y][x]-'0') - 1
			if colMap[y][n] {
				return false
			}
			if rowMap[x][n] {
				return false
			}
			if boxMap[boxId(x, y)][n] {
				return false
			}
			colMap[y][n] = true
			rowMap[x][n] = true
			boxMap[boxId(x, y)][n] = true
		}
	}
	return true
}

func newMap() [][]bool {
	m := make([][]bool, 9)
	for i, _ := range m {
		m[i] = make([]bool, 9)
	}
	return m
}

func boxId(x, y int) int {
	dx := x / 3
	dy := y / 3
	return dy*3 + dx
}

const (
	Empty = '.'
)
