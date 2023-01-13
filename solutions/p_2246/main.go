package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	parent := []int{-1, 0, 0, 1, 1, 2}
	s := "abacbe"
	fmt.Println(longestPath(parent, s))
}

func case2() {
	parent := []int{-1, 0, 0, 0}
	s := "aabc"
	fmt.Println(longestPath(parent, s))
}

func case3() {
	parent := []int{-1, 0, 1, 2, 3, 4}
	s := "zzabab"
	fmt.Println(longestPath(parent, s))
}
