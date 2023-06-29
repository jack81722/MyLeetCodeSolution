package main

import (
	"container/heap"
)

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	graph := map[int]map[int]float64{}
	for i, edge := range edges {
		if graph[edge[0]] == nil {
			graph[edge[0]] = map[int]float64{}
		}
		graph[edge[0]][edge[1]] = succProb[i]
		if graph[edge[1]] == nil {
			graph[edge[1]] = map[int]float64{}
		}
		graph[edge[1]][edge[0]] = succProb[i]
	}
	h := make(ProbHeap, 0, 100)
	visit := make([]bool, n)
	probs := make([]float64, n)
	heap.Init(&h)
	heap.Push(&h, &Vertex{start, 1})
	for h.Len() > 0 {
		cur := heap.Pop(&h).(*Vertex)
		visit[cur.node] = true
		for next, prob := range graph[cur.node] {
			if visit[next] {
				continue
			}
			newProb := prob * cur.probFromStart
			if probs[next] >= newProb {
				continue
			}
			probs[next] = newProb
			heap.Push(&h, &Vertex{next, newProb})
		}
	}
	return probs[end]
}

type Vertex struct {
	node          int
	probFromStart float64
}

type ProbHeap []*Vertex

func (p ProbHeap) Len() int {
	return len(p)
}

func (p ProbHeap) Less(i, j int) bool { return p[i].probFromStart > p[j].probFromStart }

func (p ProbHeap) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p *ProbHeap) Push(x interface{}) {
	*p = append(*p, x.(*Vertex))
}

func (p *ProbHeap) Pop() interface{} {
	n := p.Len()
	x := (*p)[n-1]
	*p = (*p)[:n-1]
	return x
}
