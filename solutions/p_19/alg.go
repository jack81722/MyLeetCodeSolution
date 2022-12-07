package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	list := make([]*ListNode, 0, 1000)

	cur := head
	list = append(list, cur)
	for cur.Next != nil {
		cur = cur.Next
		list = append(list, cur)
	}
	end := len(list)
	if end <= end-n+1 {
		list[end-n-1].Next = nil
		return head
	}
	if end-n-1 < 0 {
		head = list[end-n+1]
		return head
	}
	list[end-n-1].Next = list[end-n+1]
	return head
}
