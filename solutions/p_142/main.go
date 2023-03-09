package main

import (
	"leetcode/scenarios"
	"log"
)

func main() {
	case1()
}

func case1() {
	pos := 1
	arr := []int{3, 2, 0, -4}
	head := createList(pos, arr...)
	log.Println(detectCycle(head).Val)
}

func createList(pos int, n ...int) (head *scenarios.ListNode) {
	cur := scenarios.NewList(n[0])
	head = cur
	var cycleNode *scenarios.ListNode
	for i := 1; i < len(n); i++ {
		cur.Next = scenarios.NewList(n[i])
		cur = cur.Next
		if pos == i {
			cycleNode = cur
		}
	}
	cur.Next = cycleNode
	return
}
