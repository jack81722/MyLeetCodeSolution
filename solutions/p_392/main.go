package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	str1 := "abc"
	str2 := "ahbgdc"
	fmt.Println(isSubsequence(str1, str2))
}

func case2() {
	str1 := "axc"
	str2 := "ahbgdc"
	fmt.Println(isSubsequence(str1, str2))
}
