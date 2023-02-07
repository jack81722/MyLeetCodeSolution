package main

func totalFruit(fruits []int) int {
	m := 0
	f1 := fruits[0]
	f2 := fruits[0]
	c := 0
	last := 0
	for i := 0; i < len(fruits); i++ {
		f := fruits[i]
		if f2 != f {
			if f1 != f && f1 != f2 {
				m = max(m, c)
				c = i - last
			}
			f1 = f2
			f2 = f
			last = i
		}
		c++
	}
	return max(m, c)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
