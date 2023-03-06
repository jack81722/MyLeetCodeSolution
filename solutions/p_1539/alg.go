package main

func findKthPositive(arr []int, k int) int {
	n := 1
	i := 0
	for ; k > 0; n++ {
		if i < len(arr) && n == arr[i] {
			i++
		} else {
			k--
		}
	}
	return n - 1
}
