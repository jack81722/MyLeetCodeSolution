package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	n := 4
	roads := [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}
	fmt.Println(minScore(n, roads))
}

func case2() {
	n := 4
	roads := [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}
	fmt.Println(minScore(n, roads))
}
