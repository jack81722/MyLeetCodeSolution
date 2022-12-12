package main

func addDigits(num int) int {
	sum := num
	for sum >= 10 {
		arr := toList(sum)
		sum = 0
		for _, n := range arr {
			sum += n
		}
	}
	return sum
}

func toList(n int) []int {
	arr := make([]int, 0, 10)
	for {
		arr = append(arr, n%10)
		n /= 10
		if n == 0 {
			break
		}
	}
	return arr
}
