package main

import "fmt"

func main() {
	case1()
}

func case1() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	fmt.Println(buildTree(inorder, postorder))
}
