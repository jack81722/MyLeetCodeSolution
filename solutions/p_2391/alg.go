package main

func garbageCollection(garbage []string, travel []int) int {
	p, m, g := 0, 0, 0
	lastP, lastM, lastG := 0, 0, 0
	travel = append([]int{0}, travel...)
	for i := 0; i < len(travel); i++ {
		for j := range garbage[i] {
			switch garbage[i][j] {
			case P:
				p++
				lastP = i
			case M:
				m++
				lastM = i
			case G:
				g++
				lastG = i
			}
		}
		if i > 0 {
			travel[i] += travel[i-1]
		}
	}
	return p + m + g + travel[lastP] + travel[lastM] + travel[lastG]
}

const (
	P = 'P'
	M = 'M'
	G = 'G'
)
