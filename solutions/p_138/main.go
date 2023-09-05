package main

import "fmt"

func main() {
	case1()
}

func case1() {
	n0 := &Node{Val: 7}
	n1 := &Node{Val: 13}
	n2 := &Node{Val: 11}
	n3 := &Node{Val: 10}
	n4 := &Node{Val: 1}
	n0.Next = n1
	n0.Random = nil
	n1.Next = n2
	n1.Random = n0
	n2.Next = n3
	n2.Random = n4
	n3.Next = n4
	n3.Random = n2
	n4.Next = nil
	n4.Random = n1
	cloned := copyRandomList(n0)
	cur := cloned
	for cur != nil {
		fmt.Println("cur:", cur)
		if cur.Random != nil {
			fmt.Println("rand:", cur.Random)
		}
		cur = cur.Next
	}
}
