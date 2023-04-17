package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	results := make([]bool, len(candies))
	for i, candy := range candies {
		results[i] = candy+extraCandies >= max
	}
	return results
}
