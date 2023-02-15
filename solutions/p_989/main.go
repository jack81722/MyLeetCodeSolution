package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	num := []int{1, 2, 0, 0}
	k := 34
	fmt.Println(addToArrayForm(num, k))
}

func case2() {
	num := []int{2, 7, 4}
	k := 181
	fmt.Println(addToArrayForm(num, k))
}

func case3() {
	num := []int{2, 1, 5}
	k := 806
	fmt.Println(addToArrayForm(num, k))
}

func case4() {
	num := []int{0}
	k := 23
	fmt.Println(addToArrayForm(num, k))
}
