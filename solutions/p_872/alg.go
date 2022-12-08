package main

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	stack1 := NewStack(100)
	stack2 := NewStack(100)
	stack1.Push(root1)
	stack2.Push(root2)
	arr1, arr2 := make([]int, 0, 100), make([]int, 0, 100)
	for !stack1.Empty() && !stack2.Empty() {
		for !stack1.Empty() && len(arr1) <= len(arr2) {
			arr1 = traversalStep(stack1, arr1)
		}
		for !stack2.Empty() && len(arr2) < len(arr1) {
			arr2 = traversalStep(stack2, arr2)
		}
		if len(arr1) > 0 && len(arr1) == len(arr2) && arr1[len(arr1)-1] != arr2[len(arr2)-1] {
			return false
		}
	}
	for !stack1.Empty() && len(arr1) <= len(arr2) {
		arr1 = traversalStep(stack1, arr1)
		if len(arr1) == len(arr2) {
			if arr1[len(arr1)-1] != arr2[len(arr2)-1] {
				return false
			}
		}
	}
	for !stack2.Empty() && len(arr2) < len(arr1) {
		arr2 = traversalStep(stack2, arr2)
		if len(arr1) == len(arr2) {
			if arr1[len(arr1)-1] != arr2[len(arr2)-1] {
				return false
			}
		}
	}
	return len(arr1) == len(arr2) && stack1.Empty() && stack2.Empty() && arr1[len(arr1)-1] == arr2[len(arr2)-1]
}

func traversalStep(stack *Stack, queue []int) []int {
	current := stack.Pop()
	if current.Left == nil && current.Right == nil {
		// is leaf
		queue = append(queue, current.Val)
	}
	if current.Right != nil {
		stack.Push(current.Right)
	}
	if current.Left != nil {
		stack.Push(current.Left)
	}
	return queue
}

type Stack struct {
	arr   []*TreeNode
	count int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		arr:   make([]*TreeNode, 0, capacity),
		count: 0,
	}
}

func (s *Stack) Empty() bool {
	return s.count < 1
}

func (s *Stack) Push(n *TreeNode) {
	if len(s.arr) > s.count {
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
