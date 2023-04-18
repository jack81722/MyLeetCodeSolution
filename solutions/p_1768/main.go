package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	word1 := "abc"
	word2 := "pqr"
	fmt.Println(mergeAlternately(word1, word2))
}

func case2() {
	word1 := "ab"
	word2 := "pqrs"
	fmt.Println(mergeAlternately(word1, word2))
}

func case3() {
	word1 := "abcd"
	word2 := "pq"
	fmt.Println(mergeAlternately(word1, word2))
}
