package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur := head
	list := make([]*ListNode, 0, 1000)
	for cur != nil {
		list = append(list, cur)
		cur = cur.Next
	}
	temp := make([]*ListNode, 0, len(list))
	i := 0
	for i = 0; i+k-1 < len(list); i += k {
		reversed := reverseList(list[i : i+k])
		temp = append(temp, reversed...)
	}
	for i = len(temp); i < len(list); i++ {
		temp = append(temp, list[i])
	}
	result := temp[0]
	cur = result
	for i = 1; i < len(temp); i++ {
		cur.Next = temp[i]
		cur = cur.Next
	}
	cur.Next = nil
	return result
}

func reverseList(list []*ListNode) []*ListNode {
	result := make([]*ListNode, len(list))
	for i, n := range list {
		result[len(list)-i-1] = n
	}
	return result
}
