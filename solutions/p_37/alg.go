package main

func solveSudoku(board [][]byte) {
	l := len(board)
	// init maps
	rowMap := newMap(l)
	colMap := newMap(l)
	boxMap := newMap(l)
	for y, _ := range board {
		for x, _ := range board[y] {
			if isEmpty(board[y][x]) {
				continue
			}
			n := int(board[y][x]-Zero) - 1
			rowMap[y][n] = true
			colMap[x][n] = true
			boxMap[boxId(x, y)][n] = true
		}
	}
	// first generation
	next := 0
	x, y := toPoint(next)
	for !isEmpty(board[y][x]) && next < Max*Max {
		next++
		x, y = toPoint(next)
	}
	for _, c := range candidateInBox(x, y, boxMap) {
		if isDuplicate(c, x, y, rowMap, colMap) {
			continue
		}
		board[y][x] = byte(c + Zero + 1)
		rowMap[y][c] = true
		colMap[x][c] = true
		boxMap[boxId(x, y)][c] = true
		if finish := solve(board, c, next, rowMap, colMap, boxMap); finish {
			return
		}
	}
}

func solve(board [][]byte, n, index int, rowMap, colMap, boxMap [][]bool) (finish bool) {
	x, y := toPoint(index)
	// skip if next not empty
	next := index + 1
	x, y = toPoint(next)
	if next >= Max*Max {
		return true
	}
	for !isEmpty(board[y][x]) && next < Max*Max {
		next++
		x, y = toPoint(next)
		if next >= Max*Max {
			return true
		}
	}
	// branches of candidate
	for _, c := range candidateInBox(x, y, boxMap) {
		if isDuplicate(c, x, y, rowMap, colMap) {
			continue
		}
		board[y][x] = byte(c + Zero + 1)
		rowMap[y][c] = true
		colMap[x][c] = true
		boxMap[boxId(x, y)][c] = true
		if finish = solve(board, c, next, rowMap, colMap, boxMap); finish {
			return
		}
	}
	// recover
	x, y = toPoint(index)
	board[y][x] = '.'
	rowMap[y][n] = false
	colMap[x][n] = false
	boxMap[boxId(x, y)][n] = false
	return false
}

func newMap(l int) [][]bool {
	m := make([][]bool, l)
	for i, _ := range m {
		m[i] = make([]bool, l)
	}
	return m
}

func boxId(x, y int) int {
	dx := x / 3
	dy := y / 3
	return dy*3 + dx
}

func isEmpty(b byte) bool {
	return b == Empty
}

func isDuplicate(n, x, y int, rowMap, colMap [][]bool) bool {
	return rowMap[y][n] || colMap[x][n]
}

func toPoint(index int) (x, y int) {
	x = index % Max
	y = index / Max
	return
}

const (
	Zero  = '0'
	Empty = '.'
	Max   = 9
)

func candidateInBox(x, y int, boxMap [][]bool) []int {
	candidate := make([]int, 0, Max)
	for n, b := range boxMap[boxId(x, y)] {
		if !b {
			candidate = append(candidate, n)
		}
	}
	return candidate
}
