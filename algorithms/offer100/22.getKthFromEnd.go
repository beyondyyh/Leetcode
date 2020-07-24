package offer100

import "fmt"

// 由于链表只能从Head开始遍历，如果要计算链表长度则需要遍历一次，
// 为减少时间复杂度，可以使用快慢指针法
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head

	// 快指针先移动k步，fast != nil是为了处理当 n > len(链表) 越界情况
	for k != 0 && fast != nil {
		fast = fast.Next
		k--
	}

	// if n >= len(链表)，返回头结点
	if fast == nil {
		return head
	}
	fmt.Println(fast.Val)

	// 快慢指针同时开始移动
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
