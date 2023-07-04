package main

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	idx := make([]int, 0, 100)
	letters := make([]int, 26)
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			idx = append(idx, i)
		}
		letters[s[i]-'a']++
	}
	if len(idx) == 0 {
		for _, l := range letters {
			if l >= 2 {
				return true
			}
		}
		return false
	}
	if len(idx) != 2 {
		return false
	}
	if s[idx[0]] != goal[idx[1]] || s[idx[1]] != goal[idx[0]] {
		return false
	}
	return true
}
