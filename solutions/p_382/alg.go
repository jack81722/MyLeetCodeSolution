package main

import (
	"leetcode/scenarios"
	"math/rand"
	"time"
)

/*
Reservoir sampling
Solution Reference: https://blog.csdn.net/Changxing_J/article/details/128185299
*/

type Solution struct {
	head *scenarios.ListNode
}

func Constructor(head *scenarios.ListNode) Solution {
	rand.Seed(time.Now().Unix())
	return Solution{head: head}
}

func (this *Solution) GetRandom() int {
	cur := this.head
	res := cur.Val
	cur = cur.Next
	i := int32(1)
	for cur != nil {
		i++
		r := rand.Int31n(i)
		if r < 1 {
			res = cur.Val
		}
		cur = cur.Next
	}
	return res
}
