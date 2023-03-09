package main

import "leetcode/scenarios"

func detectCycle2(head *scenarios.ListNode) *scenarios.ListNode {
	mp := map[*scenarios.ListNode]bool{}
	cur := head
	for cur != nil && !mp[cur] {
		mp[cur] = true
		cur = cur.Next
	}
	return cur
}
