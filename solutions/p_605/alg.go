package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	var canPlace int
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}
		if i-1 < 0 && i+1 >= len(flowerbed) {
			flowerbed[i] = 1
			canPlace++
		} else if i-1 < 0 {
			if flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				canPlace++
			}
		} else if i+1 >= len(flowerbed) {
			if flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				canPlace++
			}
		} else if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			canPlace++
		}
	}
	return canPlace >= n
}
