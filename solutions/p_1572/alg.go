package main

func diagonalSum(mat [][]int) int {
	sum := 0
	l := len(mat)
	for i := range mat {
		j := l - 1 - i
		if i == j {
			sum += mat[i][i]
			continue
		}
		sum += mat[i][i] + mat[i][j]
	}
	return sum
}
