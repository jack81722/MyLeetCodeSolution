package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	customers := "YYNY"
	fmt.Println(bestClosingTime(customers))
}

func case2() {
	customers := "NNNNN"
	fmt.Println(bestClosingTime(customers))
}

func case3() {
	customers := "YYYY"
	fmt.Println(bestClosingTime(customers))
}

func case4() {
	customers := "YNYY"
	fmt.Println(bestClosingTime(customers))
}
