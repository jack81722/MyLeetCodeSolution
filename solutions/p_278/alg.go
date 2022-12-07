package main

func firstBadVersion(n int) int {
	l := 1
	r := n
	for l < r {
		c := (l + r) / 2
		if isBadVersion(c) {
			r = c
		} else {
			l = c + 1
		}
	}
	return l
}
