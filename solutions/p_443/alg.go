package main

func compress(chars []byte) int {
	cur := chars[0]
	start := 0
	count := 1
	for i := 1; i < len(chars); i++ {
		if chars[i] != cur {
			chars[start] = cur
			start++
			start = writeCount(chars, start, count)
			cur = chars[i]
			count = 1
			continue
		}
		count++
	}
	chars[start] = cur
	start++
	start = writeCount(chars, start, count)
	return start
}

func writeCount(chars []byte, start, count int) int {
	if count > 1 {
		countStart := start
		for count > 0 {
			chars[start] = byte(count%10) + '0'
			count /= 10
			start++
		}
		for j := 0; j < (start-countStart)/2; j++ {
			chars[countStart+j], chars[start-j-1] = chars[start-j-1], chars[countStart+j]
		}
	}
	return start
}
