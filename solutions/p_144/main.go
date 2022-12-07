package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:  2,
			Left: &TreeNode{Val: 3},
		},
	}
	fmt.Println(preorderTraversal(root))
}
