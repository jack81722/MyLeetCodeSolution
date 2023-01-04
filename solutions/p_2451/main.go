package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	words := []string{"adc", "wzy", "abc"}
	fmt.Println(oddString(words))
}

func case2() {
	words := []string{"aaa", "bob", "ccc", "ddd"}
	fmt.Println(oddString(words))
}
