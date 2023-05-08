package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(diagonalSum(mat))
}

func case2() {
	mat := [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1},
		{1, 1, 1, 1}}
	fmt.Println(diagonalSum(mat))
}
