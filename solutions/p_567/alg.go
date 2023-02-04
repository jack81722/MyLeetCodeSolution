package main

func checkInclusion(s1 string, s2 string) bool {
	dic := [26]int{}
	for _, c := range s1 {
		dic[c-'a']++
	}
	count := [26]int{}
	for i := 0; i < len(s2); i++ {
		c := s2[i] - 'a'
		if i >= len(s1) {
			count[s2[i-len(s1)]-'a']--
		}
		count[c]++
		if count == dic {
			return true
		}
	}
	return false
}

// my solution
func checkInclusion2(s1 string, s2 string) bool {
	dic := make([]int, 26)
	for _, c := range s1 {
		dic[c-'a']++
	}
	temp := make([]int, 26)
	copy(temp, dic)
	len1, len2 := len(s1), len(s2)
	for i := 0; i < len2 && i <= len2-len1; i++ {
		// sliding window
		j := 0
		for j = 0; j < len1; j++ {
			c := s2[i+j] - 'a'
			if temp[c] == 0 {
				copy(temp, dic)
				break
			}
			temp[c]--
		}
		if j == len(s1) {
			return true
		}
	}
	return false
}
