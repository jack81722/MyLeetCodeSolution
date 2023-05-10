package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	n := 3
	fmt.Println(generateMatrix(n))
}

func case2() {
	n := 1
	fmt.Println(generateMatrix(n))
}
