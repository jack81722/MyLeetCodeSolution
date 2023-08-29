package main

type MyStack struct {
	queue *MyQueue
}

func Constructor() MyStack {
	return MyStack{
		queue: &MyQueue{},
	}
}

func (this *MyStack) Push(x int) {
	this.queue.Enqueue(x)
}

func (this *MyStack) Pop() int {
	var n int
	temp := make([]int, 0, len(*this.queue)-1)
	for len(*this.queue) > 0 {
		n = this.queue.Dequeue()
		temp = append(temp, n)
	}
	for i := 0; i < len(temp)-1; i++ {
		this.queue.Enqueue(temp[i])
	}
	return n
}

func (this *MyStack) Top() int {
	var n int
	temp := make([]int, 0, len(*this.queue)-1)
	for len(*this.queue) > 0 {
		n = this.queue.Dequeue()
		temp = append(temp, n)
	}
	for i := 0; i < len(temp); i++ {
		this.queue.Enqueue(temp[i])
	}
	return n
}

func (this *MyStack) Empty() bool {
	return len(*this.queue) < 1
}

type MyQueue []int

func (q *MyQueue) Dequeue() int {
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q *MyQueue) Enqueue(n int) {
	*q = append(*q, n)
}
