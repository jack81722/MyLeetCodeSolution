package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	rooms := [][]int{{1}, {2}, {3}, nil}
	fmt.Println(canVisitAllRooms(rooms))
}

func case2() {
	rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	fmt.Println(canVisitAllRooms(rooms))
}
