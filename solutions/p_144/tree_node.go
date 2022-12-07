package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traversal(root *TreeNode, do func(n *TreeNode)) {
	do(root)
	if root.Left != nil {
		traversal(root.Left, do)
	}
	if root.Right != nil {
		traversal(root.Right, do)
	}
}

func (t *TreeNode) String() string {
	str := ""
	do := func(n *TreeNode) {
		if len(str) < 1 {
			str = strconv.Itoa(n.Val)
			return
		}
		str += fmt.Sprintf(", %v", n.Val)
	}
	traversal(t, do)
	return str
}
