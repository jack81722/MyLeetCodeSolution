package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	s := "leet**cod*e"
	fmt.Println(removeStars(s))
}

func case2() {
	s := "erase*****"
	fmt.Println(removeStars(s))
}
