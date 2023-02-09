package main

func distinctNames(ideas []string) int64 {
	mp := make([]map[string]bool, 26)
	for _, idea := range ideas {
		c := idea[0] - 'a'
		if mp[c] == nil {
			mp[c] = map[string]bool{}
		}
		mp[c][idea[1:]] = true
	}
	same := [26][26]int64{}
	for i := 0; i < 26; i++ {
		for idea, _ := range mp[i] {
			for j := i + 1; j < 26; j++ {
				if mp[j][idea] {
					same[i][j]++
				}
			}
		}
	}
	var count int64
	for i := 0; i < 26; i++ {
		for j := i + 1; j < 26; j++ {
			count += (int64(len(mp[i])) - same[i][j]) * (int64(len(mp[j])) - same[i][j]) * 2
		}
	}
	return count
}

// brute-force
func distinctNames2(ideas []string) int64 {
	dic := map[string]bool{}
	for _, idea := range ideas {
		dic[idea] = true
	}
	var count int64
	for i := 0; i < len(ideas)-1; i++ {
		for j := i + 1; j < len(ideas); j++ {
			if i == j {
				continue
			}
			ideaA := []rune(ideas[i])
			ideaB := []rune(ideas[j])
			ideaA[0], ideaB[0] = ideaB[0], ideaA[0]
			if dic[string(ideaA)] || dic[string(ideaB)] {
				continue
			}
			count += 2
		}
	}
	return count
}
