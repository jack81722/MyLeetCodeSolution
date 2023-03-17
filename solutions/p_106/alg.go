package main

import (
	"leetcode/scenarios"
)

func buildTree(inorder []int, postorder []int) *scenarios.TreeNode {
	if len(inorder) <= 0 {
		return nil
	}
	root := &scenarios.TreeNode{
		Val: postorder[len(postorder)-1],
	}
	index := indexOf(inorder, root.Val)
	var il, ir, pl, pr []int
	il = inorder[:index]
	pl = postorder[:len(il)]
	if index+1 < len(postorder) {
		ir = inorder[index+1:]
		pr = postorder[len(il) : len(postorder)-1]
	} else {
		ir = nil
		pr = nil
	}
	l := buildTree(il, pl)
	r := buildTree(ir, pr)
	root.Left, root.Right = l, r
	return root
}

func indexOf(arr []int, key int) int {
	for i, ele := range arr {
		if ele == key {
			return i
		}
	}
	return -1
}
