package main

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	type strength struct {
		count, idx int
	}
	mp := make([]strength, len(mat))
	for i, list := range mat {
		sum := 0
		for _, n := range list {
			sum += n
		}
		mp[i] = strength{
			count: sum,
			idx:   i,
		}
	}
	sort.Slice(mp, func(i, j int) bool {
		if mp[i].count == mp[j].count {
			return mp[i].idx < mp[j].idx
		}
		return mp[i].count < mp[j].count
	})
	idx := make([]int, k)
	for i := 0; i < k; i++ {
		idx[i] = mp[i].idx
	}
	return idx
}
