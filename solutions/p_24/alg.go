package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	arr := make([]*ListNode, 0)
	for cur != nil {
		if cur.Next != nil {
			arr = append(arr, cur.Next, cur)
			cur = cur.Next.Next
		} else {
			arr = append(arr, cur)
			break
		}
	}
	result := &ListNode{
		Val: arr[0].Val,
	}
	if len(arr) < 2 {
		return result
	}
	cur = result
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{
			Val: arr[i].Val,
		}
		cur = cur.Next
	}
	return result
}
