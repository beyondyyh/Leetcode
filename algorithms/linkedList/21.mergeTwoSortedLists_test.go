package linkedList

import (
	"fmt"
	"testing"
)

// run: go test -v ListNode.go 21.*
func build_mergeTwoSortedLists_case() (*ListNode, *ListNode) {
	// 1->2->4
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = new(ListNode)
	l1.Next.Val = 2
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 4

	// 1->3->4->5
	l2 := new(ListNode)
	l2.Val = 1
	l2.Next = new(ListNode)
	l2.Next.Val = 3
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 4
	l2.Next.Next.Next = new(ListNode)
	l2.Next.Next.Next.Val = 5
	return l1, l2
}

func Test_mergeTwoSortedLists(t *testing.T) {
	l1, l2 := build_mergeTwoSortedLists_case()
	fmt.Println("l1:")
	dump(l1)
	fmt.Println("l2:")
	dump(l2)
	ln := mergeTwoSortedLists(l1, l2)
	fmt.Println("ln:")
	dump(ln)
}
