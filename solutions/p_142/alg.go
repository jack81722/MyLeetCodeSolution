package main

import (
	"leetcode/scenarios"
)

func detectCycle(head *scenarios.ListNode) *scenarios.ListNode {
	slow, fast := head, head
	var isCycle bool
	for fast != nil && fast.Next != nil {
		// log.Println(slow.Val, fast.Val)
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			isCycle = true
			break
		}
	}
	if !isCycle {
		return nil
	}
	// reference https://leetcode.com/problems/linked-list-cycle-ii/solutions/2623326/go-two-pointers-90-less-memory/
	for head != slow {
		head = head.Next
		slow = slow.Next
	}
	return head
}
