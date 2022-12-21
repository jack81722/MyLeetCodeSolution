package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	n := 4
	dislikes := [][]int{{1, 2}, {1, 3}, {2, 4}}
	fmt.Println(possibleBipartition(n, dislikes))
}

func case2() {
	n := 3
	dislikes := [][]int{{1, 2}, {1, 3}, {2, 3}}
	fmt.Println(possibleBipartition(n, dislikes))
}

func case3() {
	n := 5
	dislikes := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}
	fmt.Println(possibleBipartition(n, dislikes))
}
