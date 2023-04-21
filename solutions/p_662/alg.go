package main

import (
	"leetcode/scenarios"
)

func widthOfBinaryTree(root *scenarios.TreeNode) int {
	queue := make([]Node, 0, 100)
	temp := make([]Node, 0, 100)
	queue = append(queue, Node{node: root, index: 0})
	max := 1
	for len(queue) > 0 {
		for _, n := range queue {
			if n.node.Left != nil {
				temp = append(temp, Node{node: n.node.Left, index: n.index * 2})
			}
			if n.node.Right != nil {
				temp = append(temp, Node{node: n.node.Right, index: n.index*2 + 1})
			}
		}
		if len(temp) > 1 {
			width := temp[len(temp)-1].index - temp[0].index + 1
			if width > max {
				max = width
			}
		}
		queue = queue[:0]
		queue = append(queue, temp...)
		temp = temp[:0]
	}

	return max
}

type Node struct {
	node  *scenarios.TreeNode
	index int
}
