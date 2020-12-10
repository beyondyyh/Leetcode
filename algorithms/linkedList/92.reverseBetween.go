package linkedList

import (
	_ "fmt"

	_ "Leetcode/algorithms/kit"
)

func reverseBetween(head *ListNode, m, n int) *ListNode {
	var dummy *ListNode = &ListNode{Val: 0}
	dummy.Next = head
	pre := dummy
	for i := 1; i < m; i++ {
		pre = pre.Next
	}

	curr := pre.Next
	for i := m; i < n; i++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = pre.Next
		pre.Next = next

		// fmt.Printf("m=%d %v\n", m, kit.List2Ints(dummy.Next))
	}
	return dummy.Next
}
