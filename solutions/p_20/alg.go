package main

func isValid(s string) bool {
	stack := make([]rune, 0, 100)
	for _, b := range s {
		if b == '(' || b == '[' || b == '{' {
			stack = append(stack, b)
			continue
		}
		if b == ')' || b == ']' || b == '}' {
			if len(stack) < 1 {
				return false
			}
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if b != pair[cur] {
				return false
			}
		}
	}
	return len(stack) < 1
}

var pair = map[rune]rune{
	'(': ')',
	')': '(',
	'[': ']',
	']': '[',
	'{': '}',
	'}': '{',
}
