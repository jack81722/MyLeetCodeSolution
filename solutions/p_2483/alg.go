package main

func bestClosingTime(customers string) int {
	arr := make([]int, len(customers))
	arr[0] = penalty(customers[0])
	for i := 1; i < len(customers); i++ {
		arr[i] = arr[i-1] + penalty(customers[i])
	}
	close0 := arr[len(arr)-1]
	m := close0
	mIdx := 0
	for j := 1; j < len(arr); j++ {
		n := close0 - arr[j-1] + (j - arr[j-1])
		if n < m {
			m = n
			mIdx = j
		}
	}
	if len(arr)-close0 < m {
		return len(arr)
	}
	return mIdx
}

func penalty(r byte) int {
	if r == 'Y' {
		return 1
	}
	return 0
}
