package main

func maxSubArray(nums []int) int {
	for i := 0; i < len(nums); i++ {
		cache[i] = map[bool]int{true: NegInf, false: NegInf}
	}
	return solve(false, 0, nums)
}

func solve(must bool, index int, array []int) int {
	if index >= len(array) {
		if must {
			return 0
		}
		return NegInf
	}
	next := index + 1
	if cache[index][must] != NegInf {
		return cache[index][must]
	}
	if must {
		r := max(0, array[index]+solve(true, next, array))
		cache[index][must] = r
		return r
	}
	r := max(array[index]+solve(true, next, array), solve(false, next, array))
	cache[index][must] = r
	return r
}

func max(n ...int) int {
	maxN := NegInf
	for _, dn := range n {
		if maxN < dn {
			maxN = dn
		}
	}
	return maxN
}

var cache = map[int]map[bool]int{}

const (
	NegInf = -99999
)
