package main

import (
	. "../util/list"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := CreateList([]int{2, 9, 3, 4})
	OutputList(l1)
	l2 := CreateList([]int{9})
	OutputList(l2)
	OutputList(addTwoNumbers(l1, l2))
	OutputList(addTwoNumbers(l2, CreateList([]int{1, 9})))
}
