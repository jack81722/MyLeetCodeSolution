package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}

func case2() {
	s1 := "ab"
	s2 := "eidboaoo"
	fmt.Println(checkInclusion(s1, s2))
}

func case3() {
	s1 := "adc"
	s2 := "dcda"
	fmt.Println(checkInclusion(s1, s2))
}
