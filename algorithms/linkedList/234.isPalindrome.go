package linkedList

func isPalindrome(head *ListNode) bool {
	// 边界条件：空链表或只有一个节点的链表
	if head == nil || head.Next == nil {
		return true
	}

	var dummy *ListNode = &ListNode{Val: -1}
	dummy.Next = head
	slow, fast := dummy, dummy
	// 慢指针一次走一步，快指针一次走两步，当快指针走到终点，慢指针刚好处于中点位置
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 将fast指针置于下半段链表的起点
	fast = slow.Next
	// 断开前后两个链表
	slow.Next = nil
	// slow指针置于前半段链表的起点
	slow = dummy.Next

	// 反转后半段链表
	var pre *ListNode = nil // 保存指针前一节点的信息，用于反转
	for fast != nil {
		pre, fast, fast.Next = fast, fast.Next, pre
	}
	// pre就是反转后的fast链表

	// 前后半链表逐一比较，当链表长度为奇数时前半段链表长度比后半段多1，所以以后半段为准
	for pre != nil {
		if slow.Val != pre.Val {
			return false
		}
		pre = pre.Next
		slow = slow.Next
	}
	return true
}
