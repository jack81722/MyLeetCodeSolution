package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1
	fmt.Println(canPlaceFlowers(flowerbed, n))
}

func case2() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 2
	fmt.Println(canPlaceFlowers(flowerbed, n))
}
