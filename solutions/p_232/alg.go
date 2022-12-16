package main

type MyQueue struct {
	stack *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		stack: NewStack(),
	}
}

func (this *MyQueue) Push(x int) {
	this.stack.Push(x)
}

func (this *MyQueue) Pop() int {
	count := this.stack.Count() - 1
	temp := make([]int, 0, count)
	for len(temp) < count {
		temp = append(temp, this.stack.Pop())
	}
	n := this.stack.Pop()
	for i := len(temp) - 1; i >= 0; i-- {
		this.stack.Push(temp[i])
	}

	return n
}

func (this *MyQueue) Peek() int {
	count := this.stack.Count()
	temp := make([]int, 0, count)
	for len(temp) < count {
		temp = append(temp, this.stack.Pop())
	}
	n := temp[count-1]
	for i := len(temp) - 1; i >= 0; i-- {
		this.stack.Push(temp[i])
	}
	return n
}

func (this *MyQueue) Empty() bool {
	return this.stack.Empty()
}

// Stack
type Stack struct {
	arr   []int
	first int
	count int
}

func NewStack() *Stack {
	return &Stack{
		arr: make([]int, 0, 100),
	}
}

func (s *Stack) Bottom() int {
	b := s.arr[s.first]
	s.first++
	return b
}

func (s *Stack) Push(x int) {
	if len(s.arr) > s.count {
		s.arr[s.count] = x
	} else {
		s.arr = append(s.arr, x)
	}
	s.count++
}

func (s *Stack) Pop() int {
	x := s.arr[s.count-1]
	s.count--
	return x
}

func (s *Stack) Empty() bool {
	return !(s.first < s.count)
}

func (s *Stack) Count() int {
	return s.count - s.first
}
