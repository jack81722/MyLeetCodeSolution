package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	arr := []int{3, 5, 1}
	fmt.Println(canMakeArithmeticProgression(arr))
}

func case2() {
	arr := []int{1, 2, 4}
	fmt.Println(canMakeArithmeticProgression(arr))
}
