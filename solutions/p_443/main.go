package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	n := compress(chars)
	fmt.Println(string(chars))
	fmt.Println(n)
}

func case2() {
	chars := []byte{'a'}
	n := compress(chars)
	fmt.Println(string(chars))
	fmt.Println(n)
}

func case3() {
	chars := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	n := compress(chars)
	fmt.Println(string(chars))
	fmt.Println(n)
}
