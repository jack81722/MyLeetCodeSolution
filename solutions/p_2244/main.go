package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	tasks := []int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}
	fmt.Println(minimumRounds(tasks))
}

func case2() {
	tasks := []int{2, 3, 3}
	fmt.Println(minimumRounds(tasks))
}

func case3() {
	tasks := []int{5, 5, 5, 5}
	fmt.Println(minimumRounds(tasks))
}
