package main

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l := createList([]int{1, 2, 3})
	output(l)
}

// input reverse
func createList(datas []int) *ListNode {
	head := new(ListNode)
	l := head
	for _, d := range datas {
		l.Val = d
		node := new(ListNode)
		node.Val = d
		l.Next = node
		l = l.Next
	}
	return head
}

func output(l *ListNode) {
	for e := l; e != nil; e = e.Next {
		fmt.Print(e.Val, " ")
	}
}
