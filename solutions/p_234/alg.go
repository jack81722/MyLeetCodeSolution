package main

func isPalindrome(head *ListNode) bool {
	stack := NewStack()
	cur := head
	for {
		stack.Push(cur.Val)
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}
	cur = head
	for {
		if stack.Pop() != cur.Val {
			return false
		}
		if cur.Next == nil {
			break
		}
		cur = cur.Next
	}
	return true
}

type Stack struct {
	index int
	arr   []int
}

func NewStack() *Stack {
	return &Stack{
		index: 0,
		arr:   make([]int, 0, 100000),
	}
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
