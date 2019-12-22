package linkedList

import (
	"testing"
)

// 1->2->3->4->2
// head=[1,2,3,4], pos=1
func build_hasCycle_case() *ListNode {
	head := new(ListNode)
	head.Val = 1
	l1 := new(ListNode)
	l1.Val = 2
	head.Next = l1
	l2 := new(ListNode)
	l2.Val = 3
	l1.Next = l2
	l3 := new(ListNode)
	l3.Val = 4
	l2.Next = l3
	l3.Next = l1
	return head
}

// run: go test -v ListNode.go 141.*
func Test_hasCycle(t *testing.T) {
	head := build_hasCycle_case()
	t.Logf("%v", hasCycle(head))
}
