package main

import (
	"container/heap"
	"sort"
)

func getOrder(tasks [][]int) []int {
	allTasks := make([]*Task, len(tasks))
	for i, t := range tasks {
		allTasks[i] = NewTask(i, t)
	}
	sort.Slice(allTasks, func(i, j int) bool {
		return allTasks[i].start < allTasks[j].start
	})
	q := make(Queue, 0, len(tasks))
	heap.Init(&q)
	tick := allTasks[0].start
	for len(allTasks) > 0 {
		if allTasks[0].start > tick {
			break
		}
		heap.Push(&q, allTasks[0])
		heap.Fix(&q, allTasks[0].index)
		allTasks = allTasks[1:]
	}
	result := make([]int, 0, len(tasks))
	for len(q) > 0 {
		t := heap.Pop(&q).(*Task)
		result = append(result, t.order)
		tick += t.process
		if len(q) == 0 && len(allTasks) > 0 && tick < allTasks[0].start {
			tick = allTasks[0].start
		}
		for len(allTasks) > 0 {
			if allTasks[0].start > tick {
				break
			}
			heap.Push(&q, allTasks[0])
			heap.Fix(&q, allTasks[0].index)
			allTasks = allTasks[1:]
		}
	}
	return result
}

type Task struct {
	start   int
	process int
	order   int
	index   int
}

func NewTask(order int, t []int) *Task {
	return &Task{
		start:   t[0],
		process: t[1],
		order:   order,
	}
}

type Queue []*Task

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	if q[i].process == q[j].process {
		return q[i].order < q[j].order
	}
	return q[i].process < q[j].process
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *Queue) Push(x interface{}) {
	n := len(*q)
	task := x.(*Task)
	task.index = n
	*q = append(*q, task)
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	return item
}
