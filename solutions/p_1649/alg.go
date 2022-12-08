package main

func createSortedArray(instructions []int) int {
	arr := make([]int, 100001)
	res := 0
	for i, x := range instructions {
		res = (res + min(get(arr, x-1), i-get(arr, x))) % mod
		update(arr, x)
	}
	return res
}

const (
	mod = 1000000007
)

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func update(arr []int, i int) {
	for i < len(arr) {
		arr[i]++
		i += i & -i
	}
}

func get(arr []int, i int) int {
	res := 0
	for i > 0 {
		res += arr[i]
		i -= i & -i
	}
	return res
}
