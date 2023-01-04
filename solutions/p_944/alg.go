package main

func minDeletionSize(strs []string) int {
	count := 0
	for i := range strs[0] {
		cur := strs[0][i]
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
