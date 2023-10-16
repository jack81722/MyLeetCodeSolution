package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	row := 3
	fmt.Println(getRow(row))
}

func case2() {
	row := 0
	fmt.Println(getRow(row))
}

func case3() {
	row := 1
	fmt.Println(getRow(row))
}

func case4() {
	row := 6
	fmt.Println(getRow(row))
}
