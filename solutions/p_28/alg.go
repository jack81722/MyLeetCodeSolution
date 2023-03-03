package main

func strStr(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)
	for i := 0; i < l1-l2+1; i++ {
		if haystack[i:i+l2] == needle {
			return i
		}
	}
	return -1
}
