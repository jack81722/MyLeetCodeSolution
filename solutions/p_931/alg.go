package main

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp := make([]int, n*n)
	for i := 0; i < n; i++ {
		doRow(i, matrix, dp)
	}
	index := (n - 1) * n
	min := dp[index]
	for i := index + 1; i < len(dp); i++ {
		if min > dp[i] {
			min = dp[i]
		}
	}
	return min
}

func doRow(row int, matrix [][]int, dp []int) {
	for i, _ := range matrix[row] {
		dp[row*len(matrix)+i] = minPath(row, i, matrix, dp)
	}
}

func minPath(row, col int, matrix [][]int, dp []int) int {
	if row == 0 {
		return matrix[row][col]
	}
	start := col
	if col-1 >= 0 {
		start = col - 1
	}
	end := col
	if col+1 < len(matrix) {
		end = col + 1
	}
	index := (row-1)*len(matrix) + start
	min := dp[index]
	for i := start + 1; i <= end; i++ {
		index = (row-1)*len(matrix) + i
		if min > dp[index] {
			min = dp[index]
		}
	}
	return min + matrix[row][col]
}
