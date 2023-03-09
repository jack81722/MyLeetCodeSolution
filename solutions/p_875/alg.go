package main

func minEatingSpeed(piles []int, h int) int {
	l := 1
	r := max(piles...)
	check := func(piles []int, speed int) int {
		h := 0
		for _, p := range piles {
			h += p / speed
			if p%speed > 0 {
				h++
			}
		}
		return h
	}
	var mid, last int
	for l <= r {
		mid = (l + r) / 2
		if check(piles, mid) > h {
			l = mid + 1
		} else {
			last = mid
			r = mid - 1
		}
	}
	return last
}

func max(n ...int) int {
	m := n[0]
	for i := 1; i < len(n); i++ {
		if m < n[i] {
			m = n[i]
		}
	}
	return m
}
