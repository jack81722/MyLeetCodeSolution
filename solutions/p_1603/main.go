package main

import "fmt"

func main() {
	case1()
}

func case1() {
	sys := Constructor(1, 1, 0)
	fmt.Println(sys.AddCar(1))
	fmt.Println(sys.AddCar(2))
	fmt.Println(sys.AddCar(3))
	fmt.Println(sys.AddCar(1))
}
