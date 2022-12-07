package main

import "math"

func sortList(head *ListNode) *ListNode {
	var tree *Tree
	cur := head
	count := 0
	for cur != nil {
		tree = Add(tree, cur.Val)
		cur = cur.Next
		count++
	}
	result := make([]*Tree, 0, count)
	result = Traversal(tree, result)
	return ToList(result)
}

// tree

type Tree struct {
	Val    int
	l      *Tree
	r      *Tree
	height int
}

func Add(root *Tree, n int) *Tree {
	if root == nil {
		return &Tree{Val: n}
	}
	if root.Val >= n {
		root.l = Add(root.l, n)
	} else {
		root.r = Add(root.r, n)
	}
	updateHeight(root)
	root = updateBalanceFactor(root)
	return root
}

func Traversal(root *Tree, result []*Tree) []*Tree {
	if root == nil {
		return result
	}
	if root.l != nil {
		result = Traversal(root.l, result)
	}
	result = append(result, root)
	if root.r != nil {
		result = Traversal(root.r, result)
	}
	return result
}

func ToList(list []*Tree) *ListNode {
	var head, cur *ListNode
	for _, e := range list {
		if head == nil {
			head = &ListNode{
				Val: e.Val,
			}
			cur = head
			continue
		}
		cur.Next = &ListNode{
			Val: e.Val,
		}
		cur = cur.Next
	}
	return head
}

func getHeight(n *Tree) int {
	if n == nil {
		return -1
	}
	return n.height
}

func updateHeight(n *Tree) {
	if n == nil {
		return
	}
	n.height = int(math.Max(float64(getHeight(n.l)), float64(getHeight(n.r)))) + 1
}

func updateBalanceFactor(n *Tree) *Tree {
	bf := getBalanceFactor(n)
	if bf > 1 {
		if getBalanceFactor(n.l) == 1 {
			return rightRotate(n)
		} else {
			n.l = leftRotate(n.l)
			return rightRotate(n)
		}
	} else if bf < -1 {
		if getBalanceFactor(n.r) == -1 {
			return leftRotate(n)
		} else {
			n.r = rightRotate(n.r)
			return leftRotate(n)
		}
	}
	return n
}

func getBalanceFactor(n *Tree) int {
	if n == nil {
		return 0
	}
	return getHeight(n.l) - getHeight(n.r)
}

func rightRotate(n *Tree) *Tree {
	l := n.l
	temp := l.r

	n.l = temp
	l.r = n

	updateHeight(n)
	updateHeight(l)
	return l
}

func leftRotate(n *Tree) *Tree {
	r := n.r
	temp := r.l

	n.r = temp
	r.l = n

	updateHeight(n)
	updateHeight(r)
	return r
}
