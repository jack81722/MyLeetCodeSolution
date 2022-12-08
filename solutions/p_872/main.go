package main

import (
	"fmt"
)

func main() {
	case1()
	case2()
	case3()
}

func case1() {
	root1 := NewTree(3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4)
	root2 := NewTree(3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8)
	fmt.Println(leafSimilar(root1, root2))
}

func case2() {
	root1 := NewTree(1, 2, 3)
	root2 := NewTree(1, 3, 2)
	fmt.Println(leafSimilar(root1, root2))
}

func case3() {
	root1 := NewTree(3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4)
	root2 := NewTree(3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 11, nil, nil, 8, 10)
	fmt.Println(leafSimilar(root1, root2))
}
