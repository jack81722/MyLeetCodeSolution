package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	n := 4
	connections := [][]int{{0, 1}, {0, 2}, {1, 2}}
	fmt.Println(makeConnected(n, connections))
}

func case2() {
	n := 6
	connections := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}
	fmt.Println(makeConnected(n, connections))
}

func case3() {
	n := 6
	connections := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}
	fmt.Println(makeConnected(n, connections))
}
