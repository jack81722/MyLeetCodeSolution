package main

import "fmt"

func main() {
	// case1()
	// case2()
	// case3()
	case4()
}

func case1() {
	tasks := [][]int{
		{1, 2},
		{2, 4},
		{3, 2},
		{4, 1},
	}
	fmt.Println(getOrder(tasks))
}

func case2() {
	tasks := [][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}
	fmt.Println(getOrder(tasks))
}

func case3() {
	tasks := [][]int{{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24}, {2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24}}
	fmt.Println(getOrder(tasks))
}

func case4() {
	tasks := [][]int{{46, 9}, {46, 42}, {30, 46}, {30, 13}, {30, 24}, {30, 5}, {30, 21}, {29, 46}, {29, 41}, {29, 18}, {29, 16}, {29, 17}, {29, 5}, {22, 15}, {22, 13}, {22, 25}, {22, 49}, {22, 44}}
	fmt.Println(getOrder(tasks))
}
