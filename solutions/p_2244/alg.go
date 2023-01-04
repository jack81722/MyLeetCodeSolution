package main

func minimumRounds(tasks []int) int {
	mp := map[int]int{}
	for _, t := range tasks {
		mp[t]++
	}
	round := 0
	for _, v := range mp {
		if v == 1 {
			return -1
		}
		round += v / 3
		if v%3 > 0 {
			round++
		}
	}
	return round
}
