package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	time := []int{1, 2, 3}
	totalTrips := 5
	fmt.Println(minimumTime(time, totalTrips))
}

func case2() {
	time := []int{2}
	totalTrips := 1
	fmt.Println(minimumTime(time, totalTrips))
}

func case3() {
	time := []int{5, 10, 10}
	totalTrips := 9
	fmt.Println(minimumTime(time, totalTrips))
}

func case4() {
	time := []int{3, 3, 8}
	totalTrips := 6
	fmt.Println(minimumTime(time, totalTrips))
}
