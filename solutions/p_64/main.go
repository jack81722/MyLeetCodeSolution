package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid))
}

func case2() {
	grid := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(minPathSum(grid))
}
