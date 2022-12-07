package main

import "fmt"

func main() {
	l1 := NewList(2, 4, 3)
	l2 := NewList(5, 6, 4)
	fmt.Println(addTwoNumbers(l1, l2))
}
