package main

func isAlienSorted(words []string, order string) bool {
	dic := make([]int, 26)
	for o, s := range order {
		dic[s-'a'] = o
	}
	maxLen := max(words...)
	closed := make([]bool, len(words))
	last := -1
	for i := 0; i < maxLen; i++ {
		if len(closed) == len(words)-1 {
			return true
		}
		for j := 0; j < len(words); j++ {
			if closed[j] {
				continue
			}
			if last == -1 {
				last = j
				continue
			}
			curVal, lastVal := getOrder(words[j], i, dic), getOrder(words[last], i, dic)
			if curVal < lastVal {
				return false
			}
			if curVal > lastVal {
				closed[j] = true
			}
			last = j
		}
		last = -1
	}
	return true
}

func max(n ...string) int {
	m := len(n[0])
	for i := 1; i < len(n); i++ {
		if m < len(n[i]) {
			m = len(n[i])
		}
	}
	return m
}

func getOrder(str string, index int, dic []int) int {
	if index >= len(str) {
		return -1
	}
	return dic[str[index]-'a']
}
