package main

func myAtoi(s string) int {
	bytes := []byte(s)
	negative := false
	result := int64(0)
	converting := false
	for _, b := range bytes {
		if result == 0 && b == zero {
			converting = true
			continue
		}
		if b == space {
			if converting {
				return int(result)
			}
			continue
		}
		if b == plus {
			if converting {
				return int(result)
			}
			converting = true
			continue
		}
		if b == minus {
			if converting {
				return int(result)
			}
			negative = true
			converting = true
			continue
		}
		if !isDigit(b) {
			return int(result)
		}

		if isDigit(b) {
			if negative {
				result = result*10 - toDigit(b)
			} else {
				result = result*10 + toDigit(b)
			}
			converting = true
		}
		if result >= MaxInt {
			return int(MaxInt)
		}
		if result <= MinInt {
			return int(MinInt)
		}
	}
	if result >= MaxInt {
		return int(MaxInt)
	}
	if result <= MinInt {
		return int(MinInt)
	}
	return int(result)
}

const (
	space  = byte(' ')
	zero   = byte('0')
	nine   = byte('9')
	plus   = byte('+')
	minus  = byte('-')
	MaxInt = int64(2147483647)
	MinInt = int64(-2147483648)
)

func isDigit(b byte) bool {
	return b >= zero && b <= nine
}

func toDigit(b byte) int64 {
	return int64(b - zero)
}
