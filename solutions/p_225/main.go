package main

import "fmt"

func main() {
	case1()
}

func case1() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	stack.Empty()
}
