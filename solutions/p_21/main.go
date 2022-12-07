package main

import "fmt"

func main() {
	l1 := NewList(1, 2, 4)
	l2 := NewList(1, 3, 4)
	fmt.Println(mergeTwoLists(l1, l2))
}
