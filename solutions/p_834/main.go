package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	n := 6
	edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}
	fmt.Println(sumOfDistancesInTree(n, edges))
}

func case2() {
	n := 1
	edges := [][]int{}
	fmt.Println(sumOfDistancesInTree(n, edges))
}

func case3() {
	n := 2
	edges := [][]int{{1, 0}}
	fmt.Println(sumOfDistancesInTree(n, edges))
}
