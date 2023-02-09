package main

import "fmt"

func main() {
	case1()
	case2()
}

func case1() {
	ideas := []string{"coffee", "donuts", "time", "toffee"}
	fmt.Println(distinctNames(ideas))
}

func case2() {
	ideas := []string{"lack", "back"}
	fmt.Println(distinctNames(ideas))
}
