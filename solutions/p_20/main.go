package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}

func case2() {
	s := "(}"
	fmt.Println(isValid(s))
}
