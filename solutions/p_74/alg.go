package main

func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1
	var i int
	for l <= r {
		i = (l + r) / 2
		num := matrix[i][0]
		if num == target {
			return true
		}
		if num > target {
			r = i - 1
		} else {
			l = i + 1
		}
	}
	if matrix[i][0] > target {
		i -= 1
	}
	if i < 0 {
		return false
	}
	l, r = 0, len(matrix[i])-1
	for l <= r {
		j := (l + r) / 2
		num := matrix[i][j]
		if num == target {
			return true
		}
		if num > target {
			r = j - 1
		} else {
			l = j + 1
		}
	}
	return false
}

// linear
func searchMatrix2(matrix [][]int, target int) bool {
	m := len(matrix)
	i := 0
	for ; i < m; i++ {
		if matrix[i][0] > target {
			break
		}
	}
	i = i - 1
	if i < 0 {
		return false
	}
	for _, num := range matrix[i] {
		if num == target {
			return true
		}
	}
	return false
}
