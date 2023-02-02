package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	words := []string{"hello", "leetcode"}
	order := "hlabcdefgijkmnopqrstuvwxyz"
	fmt.Println(isAlienSorted(words, order))
}

func case2() {
	words := []string{"word", "world", "row"}
	order := "worldabcefghijkmnpqstuvxyz"
	fmt.Println(isAlienSorted(words, order))
}

func case3() {
	words := []string{"apple", "app"}
	order := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(isAlienSorted(words, order))
}
