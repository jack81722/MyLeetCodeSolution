package main

func romanToInt(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		next := ""
		if i < len(s)-1 {
			next = string(s[i+1])
		}
		if c == "C" && i < len(s)-1 {
			if next == "D" {
				sum += 400
				i++
				continue
			}
			if next == "M" {
				sum += 900
				i++
				continue
			}
		}

		if c == "X" && i < len(s)-1 {
			if next == "L" {
				sum += 40
				i++
				continue
			}
			if next == "C" {
				sum += 90
				i++
				continue
			}
		}

		if c == "I" && i < len(s)-1 {
			if next == "V" {
				sum += 4
				i++
				continue
			}
			if next == "X" {
				sum += 9
				i++
				continue
			}
		}
		sum += Roman[c]
	}
	return sum
}

var (
	Roman = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
)
