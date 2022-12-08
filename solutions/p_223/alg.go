package main

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	aw := ax2 - ax1
	ah := ay2 - ay1
	aArea := aw * ah

	bw := bx2 - bx1
	bh := by2 - by1
	bArea := bw * bh

	xInter := (ax2 >= bx1 && bx1 >= ax1) || (bx2 > ax1 && ax1 >= bx1)
	yInter := (ay2 >= by1 && by1 >= ay1) || (by2 > ay1 && ay1 >= by1)
	iArea := 0
	if xInter && yInter {
		ix1 := max(ax1, bx1)
		ix2 := min(ax2, bx2)
		iy1 := max(ay1, by1)
		iy2 := min(ay2, by2)
		iArea = (ix2 - ix1) * (iy2 - iy1)
	}
	return aArea + bArea - iArea
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
