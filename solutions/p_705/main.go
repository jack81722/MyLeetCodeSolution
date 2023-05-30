package main

import "fmt"

func main() {
	// case1()
	case2()
}

func case1() {
	set := Constructor()
	set.Add(1)
	set.Add(2)
	fmt.Println(set.Contains(1))
	fmt.Println(set.Contains(3))
	set.Add(2)
	fmt.Println(set.Contains(2))
	set.Remove(2)
	fmt.Println(set.Contains(2))
}

func case2() {
	set := Constructor()
	set.Add(1)
	set.Add(1024 + 1)
	set.Add(2048 + 1)
	fmt.Println(set.Contains(1))
	fmt.Println(set.Contains(1024 + 1))
	fmt.Println(set.Contains(2048 + 1))
	set.Remove(1)
	fmt.Println(set.Contains(1))
	fmt.Println(set.Contains(1024 + 1))
	fmt.Println(set.Contains(2048 + 1))
}
