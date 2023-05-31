package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	sys := Constructor()
	sys.CheckIn(45, "Leyton", 3)
	sys.CheckIn(32, "Paradise", 8)
	sys.CheckIn(27, "Leyton", 10)
	sys.CheckOut(45, "Waterloo", 15)
	sys.CheckOut(27, "Waterloo", 20)
	sys.CheckOut(32, "Cambridge", 22)
	fmt.Println(sys.GetAverageTime("Paradise", "Cambridge"))
	fmt.Println(sys.GetAverageTime("Leyton", "Waterloo"))
	sys.CheckIn(10, "Leyton", 24)
	fmt.Println(sys.GetAverageTime("Leyton", "Waterloo"))
	sys.CheckOut(10, "Waterloo", 38)
	fmt.Println(sys.GetAverageTime("Leyton", "Waterloo"))
}

func case2() {
	sys := Constructor()
	sys.CheckIn(10, "Leyton", 3)
	sys.CheckOut(10, "Paradise", 8)
	fmt.Println(sys.GetAverageTime("Leyton", "Paradise"))
	sys.CheckIn(5, "Leyton", 10)
	sys.CheckOut(5, "Paradise", 16)
	fmt.Println(sys.GetAverageTime("Leyton", "Paradise"))
	sys.CheckIn(2, "Leyton", 21)
	sys.CheckOut(2, "Paradise", 30)
	fmt.Println(sys.GetAverageTime("Leyton", "Paradise"))
}

func case3() {
	sys := Constructor()
	sys.CheckIn(45, "a", 3)
	sys.CheckIn(32, "aa", 8)
	sys.CheckIn(27, "a", 10)
	sys.CheckOut(45, "ab", 15)
	sys.CheckOut(27, "ab", 20)
	sys.CheckOut(32, "b", 22)
	fmt.Println(sys.GetAverageTime("aa", "b")) // 14
	fmt.Println(sys.GetAverageTime("a", "ab")) // 11
	sys.CheckIn(10, "a", 24)
	fmt.Println(sys.GetAverageTime("a", "ab")) // 11
	sys.CheckOut(10, "ab", 38)
	fmt.Println(sys.GetAverageTime("a", "ab")) // 12
}
