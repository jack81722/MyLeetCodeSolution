package main

import (
	"math"
)

func checkStraightLine(coordinates [][]int) bool {
	v := absX([]int{coordinates[0][0] - coordinates[1][0], coordinates[0][1] - coordinates[1][1]})
	m := slope(v)
	for i := 2; i < len(coordinates); i++ {
		vv := absX([]int{coordinates[0][0] - coordinates[i][0], coordinates[0][1] - coordinates[i][1]})
		if (v[0] == 0 && vv[0] == v[0]) || (v[1] == 0 && vv[1] == v[1]) {
			continue
		}
		mm := slope(vv)
		if m != mm {
			return false
		}
	}
	return true
}

func absX(v []int) []int {
	if v[0] < 0 {
		v[0] = -v[0]
		v[1] = -v[1]
	}
	return v
}

func slope(v []int) float64 {
	var m float64
	if v[0] == 0 {
		m = float64(math.Inf(v[1]))
	} else {
		m = float64(v[1]) / float64(v[0])
	}
	return m
}
