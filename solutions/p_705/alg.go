package main

const max = 1024

type MyHashSet struct {
	table [max]*node
}

func Constructor() MyHashSet {
	return MyHashSet{
		table: [max]*node{},
	}
}

func (this *MyHashSet) Add(key int) {
	idx := key % max
	this.table[idx] = this.table[idx].Add(key)
}

func (this *MyHashSet) Remove(key int) {
	idx := key % max
	this.table[idx] = this.table[idx].Remove(key)
}

func (this *MyHashSet) Contains(key int) bool {
	idx := key % max
	return this.table[idx].Contains(key)
}

type node struct {
	val  int
	next *node
}

func (n *node) Add(val int) *node {
	if n == nil {
		return &node{val: val}
	}
	var last *node
	cur := n
	for cur != nil {
		if cur.val == val {
			return n
		}
		last = cur
		cur = cur.next
	}
	last.next = &node{val: val}
	return n
}

func (n *node) Remove(val int) *node {
	var last *node
	cur := n
	for cur != nil {
		if cur.val == val {
			if last == nil {
				return cur.next
			}
			last.next = cur.next
			return n
		}
		last = cur
		cur = cur.next
	}
	return n
}

func (n *node) Contains(val int) bool {
	cur := n
	for cur != nil {
		if cur.val == val {
			return true
		}
		cur = cur.next
	}
	return false
}
