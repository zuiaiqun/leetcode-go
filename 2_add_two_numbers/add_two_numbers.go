package main

import (
	. "../util/list"
)

/*
type ListNode struct {
	Val  int
	Next *ListNode
}

*/

// return a new list
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Len, l2Len := 0, 0
	for e := l1; e != nil; e = e.Next {
		l1Len++
	}
	for e := l2; e != nil; e = e.Next {
		l2Len++
	}
	result, bigger, smaller := new(ListNode), l1, l2
	if l2Len > l1Len {
		bigger, smaller = l2, l1
	}
	var tmp *ListNode
	left, e, f := 0, smaller, bigger
	for i := 0; e != nil; e, f = e.Next, f.Next {
		curVal := (e.Val + f.Val + left) % 10
		left = (e.Val + f.Val + left - curVal) / 10
		if i == 0 {
			tmp = result
		} else {
			tmp.Next = new(ListNode)
			tmp = tmp.Next
		}
		tmp.Val = curVal
		i++
	}
	for ; f != nil; f = f.Next {
		curVal := (f.Val + left) % 10
		left = (f.Val + left - curVal) / 10
		tmp.Next = new(ListNode)
		tmp = tmp.Next
		tmp.Val = curVal
	}
	if left > 0 {
		tmp.Next = new(ListNode)
		tmp = tmp.Next
		tmp.Val = left
	}
	return result
}
