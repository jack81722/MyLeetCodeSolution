package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
	case5()
}

func case1() {
	n := 7
	edges := [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}
	labels := "abaedcd"
	fmt.Println(countSubTrees(n, edges, labels))
}

func case2() {
	n := 4
	edges := [][]int{{0, 1}, {1, 2}, {0, 3}}
	labels := "bbbb"
	fmt.Println(countSubTrees(n, edges, labels))
}

func case3() {
	n := 5
	edges := [][]int{{0, 1}, {0, 2}, {1, 3}, {0, 4}}
	labels := "aabab"
	fmt.Println(countSubTrees(n, edges, labels))
}

func case4() {
	n := 7
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}}
	labels := "aaabaaa"
	fmt.Println(countSubTrees(n, edges, labels))
}

func case5() {
	n := 4
	edges := [][]int{{0, 2}, {0, 3}, {1, 2}}
	labels := "aeed"
	fmt.Println(countSubTrees(n, edges, labels))
}
