package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	list := CreateList(1, 2, 3)
	result := splitListToParts(list, 5)
	for _, r := range result {
		fmt.Println(ToString(r))
	}
}

func case2() {
	list := CreateList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	result := splitListToParts(list, 3)
	for _, r := range result {
		fmt.Println(ToString(r))
	}
}

func case3() {
	list := CreateList(1, 2, 3, 4, 5, 6, 7)
	result := splitListToParts(list, 3)
	for _, r := range result {
		fmt.Println(ToString(r))
	}
}
