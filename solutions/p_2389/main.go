package main

import "fmt"

func main() {
	// case1()
	// case2()
	case3()
}

func case1() {
	nums := []int{4, 5, 2, 1}
	queries := []int{3, 10, 21}
	fmt.Println(answerQueries(nums, queries))
}

func case2() {
	nums := []int{2, 3, 4, 5}
	queries := []int{1}
	fmt.Println(answerQueries(nums, queries))
}

func case3() {
	nums := []int{736411, 184882, 914641, 37925, 214915}
	queries := []int{331244, 273144, 118983, 118252, 305688, 718089, 665450}
	fmt.Println(answerQueries(nums, queries))
}
