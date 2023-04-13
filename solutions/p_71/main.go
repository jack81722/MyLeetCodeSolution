package main

import "fmt"

func main() {
	case1()
	case2()
	case3()
	case4()
	case5()
	case6()
	case7()
	case8()
	case9()
}

func case1() {
	path := "/home/"
	fmt.Println(simplifyPath(path))
}

func case2() {
	path := "/../"
	fmt.Println(simplifyPath(path))
}

func case3() {
	path := "/home//foo"
	fmt.Println(simplifyPath(path))
}

func case4() {
	path := "/a/./b/../../c/"
	fmt.Println(simplifyPath(path))
}

func case5() {
	path := "/..."
	fmt.Println(simplifyPath(path))
}

func case6() {
	path := "/..hidden"
	fmt.Println(simplifyPath(path))
}

func case7() {
	path := "/a//b////c/d//././/.."
	fmt.Println(simplifyPath(path))
}

func case8() {
	path := "/hello../world"
	fmt.Println(simplifyPath(path))
}

func case9() {
	path := "/../..ga/b/.f..d/..../e.baaeeh./.a"
	fmt.Println(simplifyPath(path))
}
