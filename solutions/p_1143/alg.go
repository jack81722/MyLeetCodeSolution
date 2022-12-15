package main

func longestCommonSubsequence(text1 string, text2 string) int {
	// init
	mp := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		mp[i] = make([]int, len(text2)+1)
	}
	//
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				mp[i+1][j+1] = mp[i][j] + 1
			} else {
				mp[i+1][j+1] = max(mp[i][j+1], mp[i+1][j])
			}
		}
	}
	return mp[len(text1)][len(text2)]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
