package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(nums ...int) *ListNode {
	if len(nums) < 1 {
		return nil
	}
	root := &ListNode{Val: nums[0]}
	cur := root
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return root
}
