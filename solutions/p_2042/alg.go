package main

func areNumbersAscending(s string) bool {
	bytes := []byte(s)
	converting := false
	last := 0
	current := 0
	for _, b := range bytes {
		if isDigit(b) {
			if converting {
				current = current*10 + toDigit(b)
			} else {
				current = toDigit(b)
				converting = true
			}
		} else {
			if converting {
				if last >= current {
					return false
				}
				last = current
				current = 0
				converting = false
			}
		}
	}
	if converting && last >= current {
		return false
	}
	return true
}

const (
	zero = byte('0')
	nine = byte('9')
)

func isDigit(b byte) bool {
	return b >= zero && b <= nine
}

func toDigit(b byte) int {
	return int(b - zero)
}
