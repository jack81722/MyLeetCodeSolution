package main

import "fmt"

func main() {
	rect1 := []int{0, 0, 2, 2}
	rect2 := []int{1, 1, 3, 3}
	fmt.Println(isRectangleOverlap(rect1, rect2))
}
