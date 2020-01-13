package linkedList

// 主要考察3个知识点，
// 1. 归并排序的整体思想: 找到链表的middle节点，然后递归对前半部分和后半部分分别进行归并排序，最后对两个以排好序的链表进行Merge
// 2. 找到一个链表的中间节点的方法: 快慢指针
// 3. 合并两个有序的链表, easy
func sortList(head *ListNode) *ListNode {
	// 回归条件，递归法需要首先确定返回条件
	if head == nil || head.Next == nil {
		return head
	}

	left, right := split(head)
	return mergeTwoSortedList(sortList(left), sortList(right))
}

func split(head *ListNode) (*ListNode, *ListNode) {
	var dummy *ListNode = &ListNode{Val: -1}
	dummy.Next = head
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 复用fast，slow
	// 将fast指针置于下半段起点
	fast = slow.Next
	// 中点断开为两个链表
	slow.Next = nil
	// slow指针置于前半段起点
	slow = dummy.Next

	// fmt.Printf("slow:%v fast:%v\n", List2Ints(slow), List2Ints(fast))
	return slow, fast
}

// mergeTwoSortedList
func mergeTwoSortedList(left, right *ListNode) *ListNode {
	var dummy *ListNode = &ListNode{Val: -1}
	curr := dummy
	for left != nil && right != nil {
		if left.Val <= right.Val {
			curr.Next, left = left, left.Next
		} else {
			curr.Next, right = right, right.Next
		}
		curr = curr.Next
	}

	if left == nil {
		curr.Next = right
	}

	if right == nil {
		curr.Next = left
	}

	return dummy.Next
}
