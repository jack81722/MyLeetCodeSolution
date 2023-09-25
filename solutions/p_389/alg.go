package main

func findTheDifference(s string, t string) byte {
	mp := make([]int, 26)
	for i, b := range s {
		mp[b-'a']++
		mp[t[i]-'a']--
	}
	mp[t[len(t)-1]-'a']--
	for j, count := range mp {
		if count < 0 {
			return byte(j + 'a')
		}
	}
	return 'a'
}
