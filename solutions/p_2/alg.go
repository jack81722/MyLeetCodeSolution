package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = &ListNode{}
	cur := result
	carry := 0
	c1 := l1
	c2 := l2
	if c1 != nil && c2 != nil {
		v := carry
		if c1 != nil {
			v += c1.Val
		}
		if c2 != nil {
			v += c2.Val
		}
		digit := v % 10
		carry = v / 10
		cur.Val = digit
		c1 = c1.Next
		c2 = c2.Next
	}
	for c1 != nil || c2 != nil || carry > 0 {
		cur.Next = &ListNode{}
		cur = cur.Next
		v := carry
		if c1 != nil {
			v += c1.Val
		}
		if c2 != nil {
			v += c2.Val
		}
		digit := v % 10
		carry = v / 10
		cur.Val = digit
		if c1 != nil {
			c1 = c1.Next
		}
		if c2 != nil {
			c2 = c2.Next
		}
	}

	return result
}
