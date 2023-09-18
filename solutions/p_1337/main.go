package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	mat := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1}}
	k := 3
	fmt.Println(kWeakestRows(mat, k))
}

func case2() {
	mat := [][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0}}
	k := 2
	fmt.Println(kWeakestRows(mat, k))
}

func case3() {
	mat := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{1, 1, 1},
		{1, 1, 1},
		{0, 0, 0},
		{1, 1, 1},
		{1, 0, 0}}
	k := 6
	fmt.Println(kWeakestRows(mat, k))
}
