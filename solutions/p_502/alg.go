package main

import (
	"container/heap"
	"sort"
)

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	// prepare projects
	projList := make([]Project, len(profits))
	for i := range profits {
		projList[i] = Project{profit: profits[i], capital: capital[i]}
	}
	sort.Slice(projList, func(i, j int) bool {
		return projList[i].capital < projList[j].capital
	})
	//
	projHeap := make(ProjectHeap, 0, len(profits))
	var i, lastIdx int
	for k > 0 {
		for i = lastIdx; i < len(projList) && projList[i].capital <= w; i++ {
			heap.Push(&projHeap, projList[i])
		}
		if projHeap.Len() == 0 {
			return w
		}
		lastIdx = i
		proj := heap.Pop(&projHeap).(Project)
		w += proj.profit
		k--
	}
	return w
}

type Project struct {
	profit  int
	capital int
}

type ProjectHeap []Project

func (p ProjectHeap) Len() int { return len(p) }

func (p ProjectHeap) Less(i, j int) bool { return p[i].profit > p[j].profit }

func (p ProjectHeap) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p *ProjectHeap) Push(x interface{}) {
	*p = append(*p, x.(Project))
}

func (p *ProjectHeap) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}
