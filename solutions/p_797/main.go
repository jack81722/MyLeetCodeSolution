package main

import "fmt"

func main() {
	// case1()
	case2()
}

func case1() {
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	fmt.Println(allPathsSourceTarget(graph))
}

func case2() {
	graph := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	fmt.Println(allPathsSourceTarget(graph))
}
