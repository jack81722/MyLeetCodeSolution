package main

import "fmt"

func main() {
	head := NewList(1, 2, 3, 4, 5)
	k := 2
	fmt.Println(reverseKGroup(head, k))
}
