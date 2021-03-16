package linkedList

// 归并排序法
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

// 快速排序法
// 快排基本思路：也是分治，找到一个轴点一般是第一个，left头结点，right尾节点
// 然后遍历元素跟轴点进行比较，比轴点小则与left交换，同时left后移一位，
// 比轴点大则与right交换，同时right迁移一位，
// 最后对left和right分别递归进行上述操作。
// 递归结束条件：数组就一个元素
// func quickSortList(head *ListNode) *ListNode {
// }
