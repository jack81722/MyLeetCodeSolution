package main

func minDeletionSize(strs []string) int {
	count := 0
	for i, c := range strs[0] {
		cur := byte(c)
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < cur {
				count++
				break
			}
			cur = strs[j][i]
		}
	}
	return count
}
