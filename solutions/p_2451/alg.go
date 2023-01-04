package main

func oddString(words []string) string {
	r1 := toIntArr(words[0])
	r2 := toIntArr(words[1])
	r3 := toIntArr(words[2])
	d12 := diff(r1, r2)
	d13 := diff(r1, r3)
	d23 := diff(r2, r3)
	if !d12 && d13 {
		return words[2]
	}
	if !d13 && d23 {
		return words[1]
	}
	if !d23 && d13 {
		return words[0]
	}

	for i := 3; i < len(words); i++ {
		if diff(r1, toIntArr(words[i])) {
			return words[i]
		}
	}
	return words[0]
}

func toIntArr(word string) []int {
	arr := make([]int, len(word)-1)
	for i := 1; i < len(word); i++ {
		arr[i-1] = int(word[i] - word[0])
	}
	return arr
}

func diff(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return true
		}
	}
	return false
}
