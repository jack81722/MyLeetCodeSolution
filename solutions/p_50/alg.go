package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	nn := n / 2
	half := myPow(x, nn)
	if n%2 == 0 {
		return half * half
	}
	if n > 0 {
		return half * half * x
	}
	return half * half / x
}
