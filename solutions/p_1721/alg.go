package main

import (
	"leetcode/scenarios"
)

func swapNodes(head *scenarios.ListNode, k int) *scenarios.ListNode {
	start, end := head, head
	cur, i := head, k
	for ; cur != nil; cur = cur.Next {
		i--
		if i == 0 {
			start = cur
		}
		if i < 0 {
			end = end.Next
		}
	}
	start.Val, end.Val = end.Val, start.Val
	return head
}
