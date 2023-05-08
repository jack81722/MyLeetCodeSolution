package main

func maxVowels(s string, k int) int {
	mp := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	var m int
	count := 0
	for i := 0; i < k; i++ {
		if mp[s[i]] {
			count++
		}
	}
	m = count
	for i := 1; i <= len(s)-k; i++ {
		if mp[s[i-1]] {
			count--
		}
		if mp[s[i+k-1]] {
			count++
		}
		if count > m {
			m = count
		}
		if m == k {
			return m
		}
	}
	return m
}
