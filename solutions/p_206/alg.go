package main

func reverseList(head *ListNode) *ListNode {
	stack := NewStack()
	for cur := head; cur != nil; cur = cur.Next {
		stack.Push(cur.Val)
	}
	if !stack.HasNext() {
		return nil
	}
	result := &ListNode{
		Val: stack.Pop(),
	}
	cur := result
	for stack.HasNext() {
		v := stack.Pop()
		cur.Next = &ListNode{
			Val: v,
		}
		cur = cur.Next
	}
	return result
}

type Stack struct {
	index int
	arr   []int
}

func NewStack() *Stack {
	return &Stack{
		index: 0,
		arr:   make([]int, 0, 1000),
	}
}

func (s *Stack) HasNext() bool {
	return s.index > 0
}

func (s *Stack) Push(v int) {
	s.arr = append(s.arr, v)
	s.index++
}

func (s *Stack) Pop() int {
	result := s.arr[s.index-1]
	s.index--
	return result
}
