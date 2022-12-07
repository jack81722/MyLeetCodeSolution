package main

func preorderTraversal(root *TreeNode) []int {
	cap := 100
	result := make([]int, 0, cap)
	if root == nil {
		return result
	}
	stack := NewStack(cap)
	stack.Push(root)
	for stack.Count() > 0 {
		cur := stack.Pop()
		result = append(result, cur.Val)
		if cur.Right != nil {
			stack.Push(cur.Right)
		}
		if cur.Left != nil {
			stack.Push(cur.Left)
		}

	}
	return result
}

type Stack struct {
	count int
	arr   []*TreeNode
}

func NewStack(cap int) *Stack {
	return &Stack{
		count: 0,
		arr:   make([]*TreeNode, 0, cap),
	}
}

func (s *Stack) Count() int {
	return s.count
}

func (s *Stack) Peek() *TreeNode {
	return s.arr[s.count-1]
}

func (s *Stack) Push(n *TreeNode) {
	if len(s.arr) <= s.count {
		s.arr = append(s.arr, n)
	} else {
		s.arr[s.count] = n
	}
	s.count++
}

func (s *Stack) Pop() *TreeNode {
	item := s.arr[s.count-1]
	s.count--
	return item
}
