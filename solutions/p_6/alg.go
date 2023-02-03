package main

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	result := make([]byte, 0, len(s))
	cycle := numRows<<1 - 2
	for r := 0; r < numRows; r++ {
		c := r
		for ; c < len(s); c += cycle {
			result = append(result, s[c])
			bonus := c + cycle - r<<1
			if r != 0 && r != numRows-1 && bonus < len(s) && bonus > 0 {
				result = append(result, s[bonus])
			}
		}
	}
	return string(result)
}
