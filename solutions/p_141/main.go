package main

import "fmt"

func main() {
	head := NewList(3, 2, 0, -4)
	doLoop(head, 1)
	fmt.Println(hasCycle(head))
}

func doLoop(head *ListNode, pos int) {
	cur := head
	var target *ListNode
	index := 0
	for cur.Next != nil {
		cur = cur.Next
		if pos == index {
			target = cur
		}
		index++
	}
	cur.Next = target
}
