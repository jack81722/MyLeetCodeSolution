package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	n := 3
	edges := [][]int{{0, 1}, {0, 2}, {1, 2}}
	fmt.Println(countPairs(n, edges))
}

func case2() {
	n := 7
	edges := [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}}
	fmt.Println(countPairs(n, edges))
}
