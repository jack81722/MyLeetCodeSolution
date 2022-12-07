package main

func checkString(s string) bool {
	bytes := []byte(s)
	ia := len(s) - 1
	ib := 0
	maxA := -1
	minB := len(s)
	for ia >= 0 || ib < len(s) {
		if bytes[ia] == a && ia > maxA {
			maxA = ia
		}
		if bytes[ib] == b && ib < minB {
			minB = ib
		}
		if maxA > minB {
			return false
		}
		ia--
		ib++
	}
	return true
}

const (
	a = byte('a')
	b = byte('b')
)
