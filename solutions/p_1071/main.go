package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
	case5()
	case6()
	case7()
	case8()
}

func case1() {
	str1 := "ABCABC"
	str2 := "ABC"
	fmt.Println(gcdOfStrings(str1, str2))
}

func case2() {
	str1 := "ABCCBA"
	str2 := "ABCCBAABCCBA"
	fmt.Println(gcdOfStrings(str1, str2))
}

func case3() {
	str1 := "ABCDEF"
	str2 := "ABC"
	fmt.Println(gcdOfStrings(str1, str2))
}

func case4() {
	str1 := "TAUXXTAUXXTAUXXTAUXXTAUXX"
	str2 := "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"
	fmt.Println(gcdOfStrings(str1, str2))
}

func case5() {
	str1 := "ABABABAB"
	str2 := "ABAB"
	fmt.Println(gcdOfStrings(str1, str2))
}

func case6() {
	str1 := "ABABC"
	str2 := "ABABCABABC"
	fmt.Println(gcdOfStrings(str1, str2))
}

func case7() {
	str1 := "LEET"
	str2 := "CODE"
	fmt.Println(gcdOfStrings(str1, str2))
}

func case8() {
	str1 := "AAAAAAAAA"
	str2 := "AAACCC"
	fmt.Println(gcdOfStrings(str1, str2))
}
