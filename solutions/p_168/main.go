package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	num := 1
	fmt.Println(convertToTitle(num))
}

func case2() {
	num := 28
	fmt.Println(convertToTitle(num))
}

func case3() {
	num := 701
	fmt.Println(convertToTitle(num))
}
