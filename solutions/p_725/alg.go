package main

func splitListToParts(head *ListNode, k int) []*ListNode {
	arr := make([]*ListNode, 0, 100)
	cur := head
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	// split
	result := make([]*ListNode, 0, 100)
	if k >= len(arr) {
		for i := 0; i < k; i++ {
			if i >= len(arr) {
				result = append(result, nil)
				continue
			}
			arr[i].Next = nil
			result = append(result, arr[i])
		}
	} else {
		seg := len(arr) / k
		mod := len(arr) % k
		for i := 0; i < len(arr); i += seg {
			if i > 0 {
				arr[i-1].Next = nil
			}
			result = append(result, arr[i])
			if mod > 0 {
				i++
				mod--
			}
		}
	}
	return result
}
