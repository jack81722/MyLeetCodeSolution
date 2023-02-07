package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	fruits := []int{1, 2, 1}
	fmt.Println(totalFruit(fruits))
}

func case2() {
	fruits := []int{0, 1, 2, 2}
	fmt.Println(totalFruit(fruits))
}

func case3() {
	fruits := []int{1, 2, 3, 2, 2}
	fmt.Println(totalFruit(fruits))
}

func case4() {
	fruits := []int{1, 0, 1, 4, 1, 4, 1, 2, 3}
	fmt.Println(totalFruit(fruits))
}
