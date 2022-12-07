package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	sp := NewStack()
	sp.Push(p)
	sq := NewStack()
	sq.Push(q)
	for sp.Count() > 0 && sq.Count() > 0 {
		// list.Add(stack.Pop())
		dp := sp.Pop()
		dq := sq.Pop()
		if dp == nil && dq == nil {
			// true
		} else if dp == nil || dq == nil {
			return false
		} else if dp.Val != dq.Val {
			return false
		}
		doPush(sp, dp)
		doPush(sq, dq)
	}
	if sp.Count() != sq.Count() {
		return false
	}
	return true
}

func doPush(stack *Stack, n *TreeNode) {
	if n == nil {
		return
	}
	stack.Push(n.Right)
	stack.Push(n.Left)
}

type Stack struct {
	arr   []*TreeNode
	count int
}

func NewStack() *Stack {
	return &Stack{
		arr:   make([]*TreeNode, 50),
		count: 0,
	}
}

func (s *Stack) Count() int {
	return s.count
}

func (s *Stack) Push(n *TreeNode) {
	if len(s.arr) >= s.count {
		s.arr[s.count] = n
	} else {
		s.arr = append(s.arr, n)
	}
	s.count++
	return
}

func (s *Stack) Pop() *TreeNode {
	n := s.arr[s.count-1]
	s.count--
	return n
}
