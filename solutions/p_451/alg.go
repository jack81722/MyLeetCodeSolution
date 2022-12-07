package main

import (
	"sort"
	"strings"
)

func frequencySort(s string) string {
	mapped := map[int32]int{}
	for _, c := range s {
		mapped[c] += 1
	}
	list := make([]int32, 0, len(mapped))
	for k, _ := range mapped {
		list = append(list, k)
	}
	sort.Slice(list, func(i, j int) bool {
		return mapped[list[i]] > mapped[list[j]]
	})
	result := ""
	for _, c := range list {
		result += strings.Repeat(string(c), mapped[c])
	}
	return result
}
