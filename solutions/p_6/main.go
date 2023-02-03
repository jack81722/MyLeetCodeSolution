package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert(s, numRows))
}

func case2() {
	s := "PAYPALISHIRING"
	numRows := 4
	fmt.Println(convert(s, numRows))
}

func case3() {
	s := "A"
	numRows := 1
	fmt.Println(convert(s, numRows))
}

func case4() {
	s := "ABCDE"
	numRows := 4
	fmt.Println(convert(s, numRows))
}
