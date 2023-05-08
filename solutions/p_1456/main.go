package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	s := "abciiidef"
	k := 3
	fmt.Println(maxVowels(s, k))
}

func case2() {
	s := "leetcode"
	k := 3
	fmt.Println(maxVowels(s, k))
}
