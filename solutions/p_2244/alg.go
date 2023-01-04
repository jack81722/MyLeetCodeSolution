package main

func minimumRounds(tasks []int) int {
	mp := map[int]int{}
	for _, t := range tasks {
		mp[t]++
	}
	round := 0
	for _, v := range mp {
		c := v / 3
		r := v % 3
		round += c
		if r == 1 {
			if c >= 1 {
				round++
				continue
			}
			return -1
		}
		if r == 2 {
			round++
		}
	}
	return round
}
