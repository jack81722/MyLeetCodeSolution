package main

import "fmt"

func main() {
	// case1()
	case2()
}

func case1() {
	n := 3
	edges := [][]int{
		{0, 1},
		{1, 2},
		{2, 0},
	}
	source := 0
	destination := 2
	fmt.Println(validPath(n, edges, source, destination))
}

func case2() {
	n := 6
	edges := [][]int{
		{0, 1},
		{0, 2},
		{3, 5},
		{5, 4},
		{4, 3},
	}
	source := 0
	destination := 5
	fmt.Println(validPath(n, edges, source, destination))
}
