package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	piles := []int{5, 4, 9}
	k := 2
	fmt.Println(minStoneSum(piles, k))
}

func case2() {
	piles := []int{4, 3, 6, 7}
	k := 3
	fmt.Println(minStoneSum(piles, k))
}

func case3() {
	piles := []int{8011, 8387, 6007, 1235, 5595, 138, 3136, 1740, 963, 9412, 802, 4475, 9695, 3713, 1742, 8559, 2237, 4356, 4012, 3449, 990, 6930, 523, 931, 5937, 5902, 2817, 4088, 385, 1359, 1888, 1106, 416, 670, 347, 6113, 4190, 2094, 2575, 3011, 8571, 32, 6318, 9658, 708, 4061, 2481, 595, 69}
	k := 93
	// 29912
	fmt.Println(minStoneSum(piles, k))
}
