package main

import "fmt"

func main() {
	case1()
}

func case1() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}
