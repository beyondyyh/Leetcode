package linkedList

func mergeTwoSortedLists(l1, l2 *ListNode) *ListNode {
	var prehead *ListNode = new(ListNode)
	cur := prehead

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 == nil {
		cur.Next = l2
	}
	if l2 == nil {
		cur.Next = l1
	}

	return prehead.Next
}
