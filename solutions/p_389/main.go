package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	s := "abcd"
	t := "abcde"
	fmt.Println(string(findTheDifference(s, t)))
}

func case2() {
	s := ""
	t := "y"
	fmt.Println(string(findTheDifference(s, t)))
}
