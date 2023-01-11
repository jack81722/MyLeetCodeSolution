package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	n := 7
	edges := [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}
	hasApple := []bool{false, false, true, false, true, true, false}
	fmt.Println(minTime(n, edges, hasApple))
}

func case2() {
	n := 7
	edges := [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}
	hasApple := []bool{false, false, true, false, false, true, false}
	fmt.Println(minTime(n, edges, hasApple))
}

func case3() {
	n := 7
	edges := [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}
	hasApple := []bool{false, false, false, false, false, false, false}
	fmt.Println(minTime(n, edges, hasApple))
}

func case4() {
	n := 4
	edges := [][]int{{0, 2}, {0, 3}, {1, 2}}
	hasApple := []bool{false, true, false, false}
	fmt.Println(minTime(n, edges, hasApple))
}
