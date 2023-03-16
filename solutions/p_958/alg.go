package main

import (
	"fmt"
	"leetcode/scenarios"
	"log"
)

func isCompleteTree(root *scenarios.TreeNode) bool {
	queue := make([]*scenarios.TreeNode, 0, 32)
	queue = append(queue, root)
	for len(queue) > 0 {
		print(queue)
		cur := queue[0]
		queue = queue[1:]
		if cur == nil && !allNil(queue) {
			return false
		}
		if cur != nil {
			queue = append(queue, cur.Left, cur.Right)
		}

	}
	return true
}

func allNil(q []*scenarios.TreeNode) bool {
	for _, n := range q {
		if n != nil {
			return false
		}
	}
	return true
}

func print(q []*scenarios.TreeNode) {
	str := ""
	for _, n := range q {
		if n != nil {
			str += fmt.Sprintf("%v, ", n.Val)
		} else {
			str += "-,"
		}
	}
	log.Println(str)
}
