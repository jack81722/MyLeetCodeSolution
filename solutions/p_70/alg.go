package main

func climbStairs(n int) int {
	count1 := n % 2
	count2 := n / 2
	sample := count1 + count2
	sum := 0
	for count2 > 0 {
		cases := cc(sample, max(count1, count2), min(count1, count2))
		sum += cases
		count1 += 2
		count2 -= 1
		sample++
	}
	return int(sum + 1)
}

func stair(n int) int {
	if n < 1 {
		return 1
	}
	result := n
	n--
	for n > 0 {
		result *= n
		n--
	}
	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func cc(n, x, y int) int {
	if n == x {
		return 1
	}
	result := n
	n--
	for n > x {
		result *= n
		if result%y == 0 {
			result /= y
			y--
		}
		n--
	}
	if y > 1 {
		result /= stair(y)
	}
	return result
}
