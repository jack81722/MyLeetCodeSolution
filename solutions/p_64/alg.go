package main

import (
	"log"
	"math"
)

func minPathSum(grid [][]int) int {
	h := len(grid)
	w := len(grid[0])
	temp := make([]int, w*h)
	log.Printf("h:%v, w:%v", h, w)
	for i := w - 1; i >= 0; i-- {
		for j := h - 1; j >= 0; j-- {
			index := toIndex(i, j, w, h)
			temp[index] = math.MaxInt
			if j+1 < h {
				temp[index] = min(temp[index], grid[j][i]+temp[index+w])
			}
			if i+1 < w {
				temp[index] = min(temp[index], grid[j][i]+temp[index+1])
			}
			if temp[index] == math.MaxInt {
				temp[index] = grid[j][i]
			}
		}
	}
	return temp[0]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func toIndex(x, y, w, h int) int {
	return y*w + x
}
