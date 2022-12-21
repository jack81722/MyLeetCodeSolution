package main

func possibleBipartition(n int, dislikes [][]int) bool {
	group := make([]int, n) // 0: none, 1: group1, 2: group2
	mp := make(map[int][]int, n)
	for _, d := range dislikes {
		mp[d[0]-1] = append(mp[d[0]-1], d[1]-1)
		mp[d[1]-1] = append(mp[d[1]-1], d[0]-1)
	}
	for k, _ := range mp {
		if group[k] == 0 && !step(k, n, mp, group) {
			return false
		}
	}
	return true
}

func step(index, max int, mp map[int][]int, group []int) bool {
	queue := make([]int, 0, max-1)
	queue = append(queue, index)
	group[index] = 1

	for len(queue) > 0 {

		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]

			for _, n := range mp[cur] {
				if group[n] == group[cur] {
					return false
				}
				if group[n] == 0 {
					group[n] = -1 * group[cur]
					queue = append(queue, n)
				}
			}
		}

	}
	return true
}
