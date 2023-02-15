package main

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	sum := make([]byte, 0, max(i, j)+1)
	var aa, bb, s, c byte
	for i >= 0 || j >= 0 {
		aa, bb = 0, 0
		if i >= 0 {
			aa = a[i] - '0'
		}
		if j >= 0 {
			bb = b[j] - '0'
		}
		s, c = add(aa, bb, c)
		sum = append(sum, s+'0')
		i--
		j--
	}
	if c > 0 {
		sum = append(sum, c+'0')
	}
	// reverse sum
	for i := 0; i < len(sum)/2; i++ {
		sum[i], sum[len(sum)-i-1] = sum[len(sum)-i-1], sum[i]
	}
	return string(sum)
}

func add(b ...byte) (s, c byte) {
	for _, bb := range b {
		s += bb
	}
	c = s / 2
	s = s % 2
	return
}

func max(x, y int) int {
	if x < y {
		return x
	}
	return y
}
