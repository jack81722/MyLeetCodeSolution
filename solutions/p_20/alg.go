package main

func isValid(s string) bool {
	stack := NewStack(1000)
	for _, b := range s {
		ss := string(b)
		if ss == "(" || ss == "[" || ss == "{" {
			stack.Push(ss)
			continue
		}
		if ss == ")" || ss == "]" || ss == "}" {
			if !stack.HasNext() {
				return false
			}
			if p := stack.Pop(); ss != pair[p] {
				return false
			}
		}
	}
	return !stack.HasNext()
}

var pair = map[string]string{
	"(": ")",
	")": "(",
	"[": "]",
	"]": "[",
	"{": "}",
	"}": "{",
}

type Stack struct {
	count int
	array []string
}

func NewStack(cap int) *Stack {
	return &Stack{
		count: 0,
		array: make([]string, 0, cap),
	}
}

func (s *Stack) HasNext() bool {
	return s.count > 0
}

func (s *Stack) Pop() string {
	result := s.array[s.count-1]
	s.count--
	return result
}

func (s *Stack) Push(p string) {
	if len(s.array) > s.count {
		s.array[s.count] = p
		s.count++
		return
	}
	s.array = append(s.array, p)
	s.count++
}
