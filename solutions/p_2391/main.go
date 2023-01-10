package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	garbage := []string{"G", "P", "GP", "GG"}
	travel := []int{2, 4, 3}
	fmt.Println(garbageCollection(garbage, travel))
}

func case2() {
	garbage := []string{"MMM", "PGM", "GP"}
	travel := []int{3, 10}
	fmt.Println(garbageCollection(garbage, travel))
}
