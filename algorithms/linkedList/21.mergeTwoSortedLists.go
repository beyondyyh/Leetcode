package linkedList

func mergeTwoSortedLists(l1, l2 *ListNode) *ListNode {
	var dummy *ListNode = &ListNode{Val: -1}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}

	if l1 == nil {
		curr.Next = l2
	}
	if l2 == nil {
		curr.Next = l1
	}

	return dummy.Next
}
