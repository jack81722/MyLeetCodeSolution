package main

import "fmt"

func main() {
	head := NewList(1, 2, 2, 1)
	fmt.Println(isPalindrome(head))
}
