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

func NewTreeNode(n int) *TreeNode {
	return &TreeNode{Val: n}
}

func NewTreeNodes(n ...interface{}) []*TreeNode {
	arr := make([]*TreeNode, len(n))
	for i, x := range n {
		if x == nil {
			arr[i] = nil
			continue
		}
		val := x.(int)
		arr[i] = NewTreeNode(val)
	}
	return arr
}

func NewTree(n ...interface{}) *TreeNode {
	return insertLevel(n, 0)
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

func insertLevel(arr []interface{}, i int) *TreeNode {
	var root *TreeNode
	if i < len(arr) {
		if arr[i] == nil {
			return nil
		}
		if x, ok := arr[i].(int); ok {
			root = NewTreeNode(x)
			if root != nil {
				root.Left = insertLevel(arr, 2*i+1)
				root.Right = insertLevel(arr, 2*i+2)
			}
		}
	}
	return root
}
