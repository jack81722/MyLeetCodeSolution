package main

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	temp := x
	reverted := 0
	for temp > 0 {
		reverted = reverted*10 + temp%10
		temp /= 10
	}
	return reverted == x
}
