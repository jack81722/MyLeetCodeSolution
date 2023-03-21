package main

func bestHand(ranks []int, suits []byte) string {
	flush := true
	var three, pair bool
	mp := make([]int, 13)
	for i := 0; i < len(suits); i++ {
		if i+1 < len(suits) && suits[i] != suits[i+1] {
			flush = false
		}
		mp[ranks[i]-1]++
		if mp[ranks[i]-1] >= 3 {
			three = true
		}
		if mp[ranks[i]-1] == 2 {
			pair = true
		}
	}
	if flush {
		return "Flush"
	}
	if three {
		return "Three of a Kind"
	}
	if pair {
		return "Pair"
	}
	return "High Card"
}
