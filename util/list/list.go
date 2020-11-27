package list

import "fmt"

// ListNode  ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateList input reverse
func CreateList(datas []int) *ListNode {
	var head, cur *ListNode
	for i, d := range datas {
		node := new(ListNode)
		node.Val = d
		// record head
		if i == 0 {
			head = node
		} else {
			// links next
			cur.Next = node
		}
		cur = node
	}
	return head
}

// OutputList ...
func OutputList(l *ListNode) {
	for e := l; e != nil; e = e.Next {
		fmt.Print(e.Val, " ")
	}
	fmt.Println()
}
