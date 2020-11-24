package main

/**
* Definition for singly-linked list.
 * type ListNode struct {
  *     Val int
   *     Next *ListNode
* }
*/

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
	left, tmp := 0, result
	for e, f := smaller, bigger; e != nil; e, f = e.Next, f.Next {
		tmp.Val = (e.Val + f.Val + left) % 10
		left = e.Val + f.Val - tmp.Val
		tmp.Next = new(ListNode)
	}
	tmp.Next = nil

}
