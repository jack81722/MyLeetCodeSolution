package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	ranks := []int{13, 2, 3, 1, 9}
	suits := []byte{'a', 'a', 'a', 'a', 'a'}
	fmt.Println(bestHand(ranks, suits))
}

func case2() {
	ranks := []int{4, 4, 2, 4, 4}
	suits := []byte{'d', 'a', 'a', 'b', 'c'}
	fmt.Println(bestHand(ranks, suits))
}

func case3() {
	ranks := []int{10, 10, 2, 12, 9}
	suits := []byte{'a', 'b', 'c', 'a', 'd'}
	fmt.Println(bestHand(ranks, suits))
}

func case4() {
	ranks := []int{12, 12, 12, 9, 9}
	suits := []byte{'b', 'd', 'a', 'c', 'b'}
	fmt.Println(bestHand(ranks, suits))
}
