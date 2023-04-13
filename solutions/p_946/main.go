package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 2, 1}
	fmt.Println(validateStackSequences(pushed, popped))
}

func case2() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 3, 5, 1, 2}
	fmt.Println(validateStackSequences(pushed, popped))
}
