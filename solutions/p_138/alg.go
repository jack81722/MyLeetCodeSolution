package main

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	mp := map[*Node]*Node{}
	cur := head
	for cur != nil {
		mp[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		mp[cur].Next = mp[cur.Next]
		mp[cur].Random = mp[cur.Random]
		cur = cur.Next
	}
	return mp[head]
}

func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	list := make([]*Node, 0, 100)
	mp := map[*Node]int{}
	cloneList := make([]*Node, 0, 100)
	cur := head
	idx := 0
	for cur != nil {
		list = append(list, cur)
		mp[cur] = idx
		clone := &Node{Val: cur.Val}
		cloneList = append(cloneList, clone)
		cur = cur.Next
		idx++
	}
	for i, c := range cloneList {
		if i < len(cloneList)-1 {
			c.Next = cloneList[i+1]
		}
		if list[i].Random != nil {
			c.Random = cloneList[mp[list[i].Random]]
		}
	}
	return cloneList[0]
}
