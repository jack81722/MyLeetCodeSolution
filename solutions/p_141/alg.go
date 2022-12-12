package main

func hasCycle(head *ListNode) bool {
	cur := head
	if cur == nil {
		return false
	}
	closed := map[*ListNode]bool{cur: true}
	for cur.Next != nil {
		cur = cur.Next
		if closed[cur] {
			return true
		}
		closed[cur] = true
	}
	return false
}
