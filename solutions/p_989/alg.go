package main

func addToArrayForm(num []int, k int) []int {
	result := make([]int, 0, len(num)+1)
	var digit, n, sum, c int
	idx := len(num) - 1
	for idx >= 0 || k > 0 {
		n = 0
		if idx >= 0 {
			n = num[idx]
		}
		sum = k%10 + n + c
		digit = sum % 10
		c = sum / 10
		result = append(result, digit)
		idx--
		k /= 10
	}
	if c > 0 {
		result = append(result, c)
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	}
	return result
}
