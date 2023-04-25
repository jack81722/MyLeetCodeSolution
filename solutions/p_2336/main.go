package main

import "fmt"

func main() {
	case1()
}

func case1() {
	set := Constructor()
	set.AddBack(2)
	fmt.Println(set.PopSmallest())
	fmt.Println(set.PopSmallest())
	fmt.Println(set.PopSmallest())
	set.AddBack(1)
	fmt.Println(set.PopSmallest())
	fmt.Println(set.PopSmallest())
	fmt.Println(set.PopSmallest())

}
