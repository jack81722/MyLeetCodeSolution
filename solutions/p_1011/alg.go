package main

func shipWithinDays(weights []int, days int) int {
	u, b := 0, 0
	for _, w := range weights {
		u += w
		if w > b {
			b = w
		}
	}

	ship := func(cap int) int {
		sum, day := 0, 0
		for _, w := range weights {
			if sum+w <= cap {
				sum += w
			} else {
				sum = w
				day++
			}
			if day >= days {
				return day
			}
		}
		return day
	}

	for u > b {
		mid := (u + b) / 2
		if ship(mid) < days {
			u = mid
			continue
		}
		b = mid + 1
	}
	return b
}
