package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	coordinates := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
	fmt.Println(checkStraightLine(coordinates))
}

func case2() {
	coordinates := [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}
	fmt.Println(checkStraightLine(coordinates))
}

func case3() {
	coordinates := [][]int{{1, 1}, {2, 2}, {2, 0}}
	fmt.Println(checkStraightLine(coordinates))
}

func case4() {
	coordinates := [][]int{{0, 0}, {0, 1}, {0, -1}}
	fmt.Println(checkStraightLine(coordinates))
}
