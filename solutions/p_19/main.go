package main

import "fmt"

func main() {
	head := NewList(1, 2, 3, 4, 5)
	n := 2
	fmt.Println(removeNthFromEnd(head, n))
}
