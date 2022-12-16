package main

import "fmt"

func main() {
	case1()
	case2()

}

func case1() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Peek())
	fmt.Println(q.Pop())
	fmt.Println(q.Empty())
}

func case2() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	fmt.Println(q.Pop())
	q.Push(5)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}
