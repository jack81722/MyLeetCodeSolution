package main

func numberOfWays(corridor string) int {
	pos := make([]int, 0, len(corridor))
	for i, c := range corridor {
		if c == S {
			pos = append(pos, i)
		}
	}

	if len(pos)%2 == 1 || len(pos) == 0 {
		return 0
	}

	result := int64(1)
	prev := pos[1]
	for j := 2; j < len(pos); j += 2 {
		v := int64(pos[j] - prev)
		result = (result * v) % MOD
		prev = pos[j+1]
	}
	return int(result)
}

const (
	S   = 'S'
	MOD = 1000000007
)
