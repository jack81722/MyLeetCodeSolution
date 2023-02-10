package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

// 2
func case1() {
	grid := [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}
	fmt.Println(maxDistance(grid))
}

// 4
func case2() {
	grid := [][]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	fmt.Println(maxDistance(grid))
}

// 2
func case3() {
	grid := [][]int{{0, 0, 1, 1, 1}, {0, 1, 1, 0, 0}, {0, 0, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 1}}
	fmt.Println(maxDistance(grid))
}
