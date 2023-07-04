package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	s := "ab"
	goal := "ba"
	fmt.Println(buddyStrings(s, goal))
}

func case2() {
	s := "ab"
	goal := "ab"
	fmt.Println(buddyStrings(s, goal))
}

func case3() {
	s := "aa"
	goal := "aa"
	fmt.Println(buddyStrings(s, goal))
}
