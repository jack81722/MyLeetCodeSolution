package main

import (
	"container/heap"
)

func reorganizeString(s string) string {
	mp := make(map[rune]int)
	for _, c := range s {
		mp[c]++
	}
	h := make(CountHeap, 0, len(mp))
	for k, v := range mp {
		h = append(h, Count{
			Char:  k,
			Count: v,
		})
	}
	heap.Init(&h)
	result := make([]rune, 0, len(s))
	for len(result) < len(s) {
		c1 := heap.Pop(&h).(Count)
		if len(result) > 0 && c1.Char == result[len(result)-1] {
			if h.Len() < 1 {
				return ""
			}
			c2 := heap.Pop(&h).(Count)
			result = append(result, c2.Char)
			c2.Count--
			if c2.Count > 0 {
				heap.Push(&h, c2)
			}
		}
		result = append(result, c1.Char)
		c1.Count--
		if c1.Count > 0 {
			heap.Push(&h, c1)
		}
	}
	return string(result)
}

type Count struct {
	Char  rune
	Count int
}

type CountHeap []Count

func (c CountHeap) Len() int {
	return len(c)
}

func (c CountHeap) Less(i, j int) bool {
	return c[i].Count > c[j].Count
}

func (c CountHeap) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c *CountHeap) Push(v interface{}) {
	*c = append(*c, v.(Count))
}

func (c *CountHeap) Pop() interface{} {
	old := *c
	n := len(*c)
	x := old[n-1]
	*c = old[:n-1]
	return x
}
