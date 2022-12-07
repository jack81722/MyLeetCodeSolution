package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 3},
		},
	}
	fmt.Println(isSymmetric(root))
}
