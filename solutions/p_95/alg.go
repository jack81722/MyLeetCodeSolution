package main

func generateTrees(n int) []*TreeNode {
	arr := newArray(n)
	result := generate(arr, 0, n-1)
	return result
}

func insert(node *TreeNode, n *TreeNode) *TreeNode {
	if node == nil {
		return n
	}
	if node.Val > n.Val {
		node.Left = insert(node.Left, n)
		return node
	}
	if node.Val < n.Val {
		node.Right = insert(node.Right, n)
		return node
	}
	return node
}

func generate(arr []int, k, m int) []*TreeNode {
	if k == m {
		return []*TreeNode{{Val: arr[k]}}
	} else if m < k || k > m {
		return nil
	}
	result := make([]*TreeNode, 0, stair(m-k+1))
	for i := k; i <= m; i++ {
		x := generate(arr, k, i-1)
		y := generate(arr, i+1, m)

		if len(x) > 0 && len(y) > 0 {
			for _, dx := range x {
				for _, dy := range y {
					root := &TreeNode{Val: arr[i]}
					root.Left = dx
					root.Right = dy
					result = append(result, root)
				}
			}
		} else if len(x) > 0 {
			for _, dx := range x {
				root := &TreeNode{Val: arr[i]}
				root.Left = dx
				result = append(result, root)
			}
		} else {
			for _, dy := range y {
				root := &TreeNode{Val: arr[i]}
				root.Right = dy
				result = append(result, root)
			}
		}
	}
	return result
}

func permutation(arr []int, k, m int, result [][]int) [][]int {
	if k == m {
		result = append(result, arr)
		return result
	}
	for i := k; i <= m; i++ {
		seg := make([]int, len(arr))
		copy(seg, arr)
		seg = swap(seg, i, k)
		result = permutation(seg, k+1, m, result)
	}
	return result
}

func stair(n int) int {
	s := 1
	for i := 2; i <= n; i++ {
		s *= i
	}
	return s
}

func newArray(n int) []int {
	result := make([]int, n)
	for i := 1; i <= n; i++ {
		result[i-1] = i
	}
	return result
}

func swap(arr []int, i, j int) []int {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
	return arr
}
