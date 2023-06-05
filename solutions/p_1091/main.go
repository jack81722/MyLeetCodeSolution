package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	grid := [][]int{{0, 1}, {1, 0}}
	fmt.Println(shortestPathBinaryMatrix(grid))
}

func case2() {
	grid := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	fmt.Println(shortestPathBinaryMatrix(grid))
}

func case3() {
	grid := [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	fmt.Println(shortestPathBinaryMatrix(grid))
}

func case4() {
	grid := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 1}}
	fmt.Println(shortestPathBinaryMatrix(grid))
}
