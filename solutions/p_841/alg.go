package main

func canVisitAllRooms(rooms [][]int) bool {
	closed := make([]bool, len(rooms))
	queue := make([]int, 0, len(rooms))
	queue = append(queue, 0)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		closed[cur] = true
		for _, key := range rooms[cur] {
			if closed[key] {
				continue
			}
			queue = append(queue, key)
		}
	}
	return allVisit(closed)
}

func allVisit(closed []bool) bool {
	b := true
	for _, c := range closed {
		b = b && c
	}
	return b
}
