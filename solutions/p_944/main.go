package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	strs := []string{"cba", "daf", "ghi"}
	fmt.Println(minDeletionSize(strs))
}

func case2() {
	strs := []string{"a", "b"}
	fmt.Println(minDeletionSize(strs))
}

func case3() {
	strs := []string{"zyx", "wvu", "tsr"}
	fmt.Println(minDeletionSize(strs))
}

func case4() {
	strs := []string{"rrjk", "furt", "guzm"}
	fmt.Println(minDeletionSize(strs))
}
