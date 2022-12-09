package main

func maxAncestorDiff(root *TreeNode) int {
	return step(root, root.Val, root.Val)
}

func step(sub *TreeNode, max, min int) int {
	if sub == nil {
		return max - min
	}
	var maxL, maxR int
	if sub.Left != nil {
		maxL = step(sub.Left, maximum(max, sub.Left.Val), minimum(min, sub.Left.Val))
	}
	if sub.Right != nil {
		maxR = step(sub.Right, maximum(max, sub.Right.Val), minimum(min, sub.Right.Val))
	}
	result := maximum(max-min, maxL, maxR)
	return result
}

func maximum(n ...int) int {
	max := n[0]
	for i := 1; i < len(n); i++ {
		if max < n[i] {
			max = n[i]
		}
	}
	return max
}

func minimum(n ...int) int {
	min := n[0]
	for i := 1; i < len(n); i++ {
		if min > n[i] {
			min = n[i]
		}
	}
	return min
}
