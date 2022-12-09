package main

import "fmt"

func main() {
	// case1()
	// case2()
	case3()
}

func case1() {
	root := NewTree(8, 3, 10, 1, 6, nil, 14, nil, nil, 4, 7, 13)
	fmt.Println(maxAncestorDiff(root))
}

func case2() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:  0,
				Left: &TreeNode{Val: 3},
			},
		},
	}
	fmt.Println(maxAncestorDiff(root))
}

func case3() {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 5,
		},
		Right: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val:   1,
						Right: &TreeNode{Val: 3},
					},
				},
			},
		},
	}
	fmt.Println(maxAncestorDiff(root))
}
