package main

import "fmt"

func main() {
	case1()
}

func case1() {
	k := 3
	nums := []int{4, 5, 8, 2}
	kthLargest := Constructor(k, nums)
	fmt.Println(kthLargest.Add(3))
	fmt.Println(kthLargest.Add(5))
	fmt.Println(kthLargest.Add(10))
	fmt.Println(kthLargest.Add(9))
	fmt.Println(kthLargest.Add(4))
	fmt.Println(kthLargest.Add(1))
}
