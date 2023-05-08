package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	senate := "RD"
	fmt.Println(predictPartyVictory(senate))
}

func case2() {
	senate := "RDD"
	fmt.Println(predictPartyVictory(senate))
}

func case3() {
	senate := "RDRDD"
	fmt.Println(predictPartyVictory(senate))
}
