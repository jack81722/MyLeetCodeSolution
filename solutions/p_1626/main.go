package main

import (
	"fmt"
)

func main() {
	case1()
	case2()
	case3()
	case4()
	case5()
	case6()
}

// 34
func case1() {
	scores := []int{1, 3, 5, 10, 15}
	ages := []int{1, 2, 3, 4, 5}
	fmt.Println(bestTeamScore(scores, ages))
}

// 16
func case2() {
	scores := []int{4, 5, 6, 5}
	ages := []int{2, 1, 2, 1}
	fmt.Println(bestTeamScore(scores, ages))
}

// 6
func case3() {
	scores := []int{1, 2, 3, 5}
	ages := []int{8, 9, 10, 1}
	fmt.Println(bestTeamScore(scores, ages))
}

// 15
func case4() {
	scores := []int{6, 3, 2, 1, 5, 4}
	ages := []int{2, 2, 2, 2, 1, 1}
	fmt.Println(bestTeamScore(scores, ages))
}

// 27
func case5() {
	scores := []int{9, 2, 8, 8, 2}
	ages := []int{4, 1, 3, 3, 5}
	fmt.Println(bestTeamScore(scores, ages))
}

// 4003
func case6() {
	scores := []int{988, 511, 618, 880, 214, 589, 576, 744, 865, 830, 478}
	ages := []int{89, 57, 20, 93, 20, 100, 8, 18, 62, 47, 45}
	fmt.Println(bestTeamScore(scores, ages))
}
