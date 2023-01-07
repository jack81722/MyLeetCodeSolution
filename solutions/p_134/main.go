package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}

func case2() {
	gas := []int{2, 3, 4}
	cost := []int{3, 4, 3}
	fmt.Println(canCompleteCircuit(gas, cost))
}

func case3() {
	gas := []int{5, 8, 2, 8}
	cost := []int{6, 5, 6, 6}
	fmt.Println(canCompleteCircuit(gas, cost))
}
