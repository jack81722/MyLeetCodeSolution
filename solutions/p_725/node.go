package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(nums ...int) *ListNode {
	if len(nums) < 1 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head
}

func ToString(head *ListNode) string {
	if head == nil {
		return "_"
	}
	cur := head
	str := fmt.Sprintf("%v", cur.Val)
	cur = cur.Next
	for cur != nil {
		str += fmt.Sprintf(", %v", cur.Val)
		cur = cur.Next
	}
	return str
}
