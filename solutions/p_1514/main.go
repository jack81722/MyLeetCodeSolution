package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
}

func case1() {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}, {0, 2}}
	probs := []float64{0.5, 0.5, 0.2}
	start := 0
	end := 2
	fmt.Println(maxProbability(n, edges, probs, start, end))
}

func case2() {
	n := 3
	edges := [][]int{{0, 1}, {1, 2}, {0, 2}}
	probs := []float64{0.5, 0.5, 0.3}
	start := 0
	end := 2
	fmt.Println(maxProbability(n, edges, probs, start, end))
}

func case3() {
	n := 3
	edges := [][]int{{0, 1}}
	probs := []float64{0.5}
	start := 0
	end := 2
	fmt.Println(maxProbability(n, edges, probs, start, end))
}

func case4() {
	n := 7
	edges := [][]int{{0, 1}, {0, 5}, {0, 4}, {1, 2}, {4, 3}, {5, 6}, {2, 3}, {6, 3}}
	probs := []float64{1, 0.8, 0.8, 1, 0.1, 0.6, 0.1, 0.3}
	start := 0
	end := 3
	fmt.Println(maxProbability(n, edges, probs, start, end))
}
