package main

import "fmt"

func main() {
	// case1()
	// case2()
	case3()
}

func case1() {
	cap := []int{2, 3, 4, 5}
	rock := []int{1, 2, 4, 4}
	addition := 2
	fmt.Println(maximumBags(cap, rock, addition))
}

func case2() {
	cap := []int{10, 2, 2}
	rock := []int{2, 2, 0}
	addition := 100
	fmt.Println(maximumBags(cap, rock, addition))
}

func case3() {
	cap := []int{91, 54, 63, 99, 24, 45, 78}
	rock := []int{35, 32, 45, 98, 6, 1, 25}
	addition := 17
	fmt.Println(maximumBags(cap, rock, addition))
}
