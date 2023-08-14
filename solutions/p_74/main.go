package main

import "fmt"

func main() {
	// case1()
	// case2()
	// case3()
	// case4()
	case5()
}

func case1() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	fmt.Println(searchMatrix(matrix, target))
}

func case2() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 13
	fmt.Println(searchMatrix(matrix, target))
}

func case3() {
	matrix := [][]int{{1}}
	target := 0
	fmt.Println(searchMatrix(matrix, target))
}

func case4() {
	matrix := [][]int{{1}}
	target := 1
	fmt.Println(searchMatrix(matrix, target))
}

func case5() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 11
	fmt.Println(searchMatrix(matrix, target))
}
