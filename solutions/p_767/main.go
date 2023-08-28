package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	s := "aab"
	fmt.Println(reorganizeString(s))
}

func case2() {
	s := "aaab"
	fmt.Println(reorganizeString(s))
}

func case3() {
	s := "acab"
	fmt.Println(reorganizeString(s))
}
