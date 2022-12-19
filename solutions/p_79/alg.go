package main

func exist(board [][]byte, word string) bool {
	Height = len(board)
	Width = len(board[0])
	closed := make(map[int]bool, Width*Height)
	for y, _ := range board {
		for x, _ := range board[y] {
			index := toIndex(x, y)
			if word[0] != board[y][x] {
				continue
			}
			closed[index] = true
			if search(board, word, 0, x, y, closed) {
				return true
			}
			closed[index] = false
		}
	}
	return false
}

var Width int
var Height int

func search(board [][]byte, word string, charIndex int, x, y int, closed map[int]bool) bool {
	// next char
	charIndex++
	if charIndex >= len(word) {
		return true
	}
	// find neighbors
	xs, ys := getNeighbors(x, y, closed)
	if len(xs) < 1 {
		return false
	}
	// next search
	for i, _ := range xs {
		nx := xs[i]
		ny := ys[i]
		if word[charIndex] != board[ny][nx] {
			continue
		}
		index := toIndex(nx, ny)
		closed[index] = true
		if search(board, word, charIndex, nx, ny, closed) {
			return true
		}
		// recover
		closed[index] = false
	}
	return false
}

func getNeighbors(x, y int, closed map[int]bool) (xs, ys []int) {
	xs = make([]int, 0, 4)
	ys = make([]int, 0, 4)
	if x-1 >= 0 && !closed[toIndex(x-1, y)] {
		xs = append(xs, x-1)
		ys = append(ys, y)
	}
	if x+1 < Width && !closed[toIndex(x+1, y)] {
		xs = append(xs, x+1)
		ys = append(ys, y)
	}
	if y-1 >= 0 && !closed[toIndex(x, y-1)] {
		xs = append(xs, x)
		ys = append(ys, y-1)
	}
	if y+1 < Height && !closed[toIndex(x, y+1)] {
		xs = append(xs, x)
		ys = append(ys, y+1)
	}
	return
}

func toIndex(x, y int) int {
	return y*Width + x
}
