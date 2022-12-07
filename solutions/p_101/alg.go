package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lpt := root.Left
	rpt := root.Right
	if lpt == nil && rpt == nil {
		return true
	}
	lst := NewStack()
	lst.Push(lpt)
	rst := NewStack()
	rst.Push(rpt)
	for lst.Count() > 0 && rst.Count() > 0 {
		ln := lst.Pop()
		rn := rst.Pop()
		if ln == nil && rn == nil {
			// true
		} else if ln == nil || rn == nil {
			return false
		} else if ln.Val != rn.Val {
			return false
		}
		leftPush(ln, lst)
		rightPush(rn, rst)
	}
	if lst.Count() != rst.Count() {
		return false
	}
	return true
}

func leftPush(n *TreeNode, stack *Stack) {
	if n == nil {
		return
	}
	stack.Push(n.Left)
	stack.Push(n.Right)
}

func rightPush(n *TreeNode, stack *Stack) {
	if n == nil {
		return
	}
	stack.Push(n.Right)
	stack.Push(n.Left)
}

func value(n *TreeNode) int {
	if n == nil {
		return 0
	}
	return n.Val
}

type Stack struct {
	arr   []*TreeNode
	count int
}

func NewStack() *Stack {
	return &Stack{
		arr:   make([]*TreeNode, 500),
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
}

func (s *Stack) Pop() *TreeNode {
	n := s.arr[s.count-1]
	s.count--
	return n
}
