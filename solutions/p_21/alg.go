package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result, cur *ListNode
	l1 := list1
	l2 := list2
	for l1 != nil && l2 != nil {
		if cur == nil {
			result = &ListNode{}
			cur = result
		} else {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
		if l1.Val <= l2.Val {
			cur.Val = l1.Val
			l1 = l1.Next
			continue
		}
		if l1.Val > l2.Val {
			cur.Val = l2.Val
			l2 = l2.Next
			continue
		}
	}
	if l1 != nil {
		if cur == nil {
			return l1
		} else {
			cur.Next = l1
		}
	}
	if l2 != nil {
		if cur == nil {
			return l2
		} else {
			cur.Next = l2
		}
	}
	return result
}
