package main

import (
	. "../util/list"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l := CreateList([]int{1, 2, 3, 4})
	OutputList(l)
}
