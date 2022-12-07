package main

import "fmt"

func main() {
	head := NewList(1, 2, 3, 4)
	fmt.Println(swapPairs(head))
}
