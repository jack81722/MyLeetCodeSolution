package main

import (
	"fmt"
)

func main() {
	case1()
	case2()
}

func case1() {
	haystack := "sadbutsad"
	needle := "sad"
	fmt.Println(strStr(haystack, needle))
}

func case2() {
	haystack := "a"
	needle := "a"
	fmt.Println(strStr(haystack, needle))
}
