package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	n := 6
	conections := [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}
	fmt.Println(minReorder(n, conections))
}

func case2() {
	n := 5
	conections := [][]int{{1, 0}, {1, 2}, {3, 2}, {3, 4}}
	fmt.Println(minReorder(n, conections))
}

func case3() {
	n := 3
	conections := [][]int{{1, 0}, {2, 0}}
	fmt.Println(minReorder(n, conections))
}
