package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(stones))
}

func case2() {
	stones := []int{2}
	fmt.Println(lastStoneWeight(stones))
}

func case3() {
	stones := []int{4, 3, 4, 3, 2}
	fmt.Println(lastStoneWeight(stones))
}
