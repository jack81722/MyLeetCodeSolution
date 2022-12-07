package main

import "strings"

func multiply(num1 string, num2 string) string {
	subList := make([][]uint8, 200)
	for i := len(num1) - 1; i >= 0; i-- {
		c := zero
		d := zero
		stack := NewStack()
		for k := 0; k < len(num1)-i-1; k++ {
			stack.Push(zero)
		}
		for j := len(num2) - 1; j >= 0; j-- {
			x := num1[i]
			y := num2[j]
			d, c = digitMul(x, y, c)
			stack.Push(d)
		}
		if c > zero {
			stack.Push(c)
		}
		subList = append(subList, stack.ToList())
	}
	sum := "0"
	for _, sub := range subList {
		sum = allPlus(sum, string(sub))
	}
	sum = strings.TrimLeft(sum, "0")
	if len(sum) < 1 {
		sum = "0"
	}
	return sum
}

const zero = uint8('0')

func digitMul(x, y, c uint8) (digit, carry byte) {
	v := (x - zero) * (y - zero)
	digit = v%10 + zero
	carry = v/10 + zero

	digit, c = digitPlus(digit, c)
	carry, _ = digitPlus(carry, c)
	return
}

func allPlus(x, y string) string {
	i := len(x) - 1
	j := len(y) - 1
	d := zero
	c := zero
	//dd := zero
	cc := zero
	stack := NewStack()
	for i >= 0 && j >= 0 {
		dx := x[i]
		dy := y[j]
		d, cc = digitPlus(dx, c)
		d, c = digitPlus(d, dy)
		c, _ = digitPlus(c, cc)
		stack.Push(d)
		i--
		j--
	}
	for ; i >= 0; i-- {
		dx := x[i]
		d, c = digitPlus(dx, c)
		stack.Push(d)
	}
	for ; j >= 0; j-- {
		dy := y[j]
		d, c = digitPlus(dy, c)
		stack.Push(d)
	}
	if c-zero > 0 {
		stack.Push(c)
	}
	return string(stack.ToList())
}

func digitPlus(x, y uint8) (digit, carry uint8) {
	v := x + y - zero*2
	digit = v%10 + zero
	carry = v/10 + zero
	return
}

type Stack struct {
	index int
	arr   []uint8
}

func NewStack() *Stack {
	return &Stack{
		index: 0,
		arr:   make([]uint8, 0, 1000),
	}
}

func (s *Stack) Push(n uint8) {
	s.arr = append(s.arr, n)
	s.index++
}

func (s *Stack) Pop() uint8 {
	r := s.arr[s.index-1]
	s.index--
	return r
}

func (s *Stack) ToList() []uint8 {
	r := make([]uint8, s.index)
	for i := s.index - 1; i >= 0; i-- {
		r[s.index-1-i] = s.arr[i]
	}
	return r
}
