package linkedList

import (
	"testing"
)

func build_removeNthFromEnd_case() *ListNode {
	head := new(ListNode)
	head.Val = 1
	l2 := new(ListNode)
	l2.Val = 2
	l3 := new(ListNode)
	l3.Val = 3
	l4 := new(ListNode)
	l4.Val = 4
	l5 := new(ListNode)
	l5.Val = 5
	head.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	return head
}

// run: go test -v ListNode.go 19.*
func Test_removeNthFromEnd1(t *testing.T) {
	head := build_removeNthFromEnd_case()
	t.Logf("head:%v\n", display(head))
	res := removeNthFromEnd1(head, 5)
	t.Logf("res:%v\n", display(res))
}

func Test_removeNthFromEnd2(t *testing.T) {
	head := build_removeNthFromEnd_case()
	t.Logf("head:%v\n", display(head))
	res := removeNthFromEnd2(head, 5)
	t.Logf("res:%v\n", display(res))
}
