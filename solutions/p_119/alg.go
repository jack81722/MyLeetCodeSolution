package main

func getRow(rowIndex int) []int {
	res := []int{1}
	prev := 1
	for k := 1; k <= rowIndex; k++ {
		next_val := prev * (rowIndex - k + 1) / k
		res = append(res, next_val)
		prev = next_val
	}
	return res
}

func getRow2(rowIndex int) []int {
	results := make([][]int, rowIndex+2)
	results[0] = []int{1}
	results[1] = []int{1, 1}
	for i := 2; i < rowIndex+1; i++ {
		results[i] = make([]int, i+1)
		results[i][0] = 1
		results[i][i] = 1
		for j := 1; j < i; j++ {
			results[i][j] = results[i-1][j-1] + results[i-1][j]
		}
	}
	return results[rowIndex]
}
