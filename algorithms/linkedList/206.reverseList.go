package linkedList

import (
	"fmt"

	"Leetcode/algorithms/kit"
)

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		fmt.Printf("before: pre:%v cur:%v\n", kit.List2Ints(pre), kit.List2Ints(cur))
		pre, cur, cur.Next = cur, cur.Next, pre
		fmt.Printf("after: pre:%v cur:%v\n\n", kit.List2Ints(pre), kit.List2Ints(cur))
	}
	fmt.Printf("pre:%v cur:%v\n", kit.List2Ints(pre), kit.List2Ints(cur))
	return pre
}

// 执行过程：
// pre:[nil] cur:[1 2 3 nil]
// pre:[1 nil] cur:[2 3 nil]
// pre:[2 1 nil] cur:[3 nil]
// pre:[3 2 1 nil] cur:[nil]

func reverseListV2(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode = nil
	// 每次循环，都将当前节点指向它前面的节点，然后当前节点和前节点后移
	for curr != nil {
		fmt.Printf("before: pre:%v cur:%v\n", kit.List2Ints(prev), kit.List2Ints(curr))
		tmpNode := curr.Next // 临时节点，暂存当前节点的下一个节点，用于后移
		curr.Next = prev     // 将当前节点指针指向它前面的节点
		prev = curr          // 前节点的指针后移
		curr = tmpNode       // 当前节点的指针后移
		fmt.Printf("after: pre:%v cur:%v\n\n", kit.List2Ints(prev), kit.List2Ints(curr))
	}
	return prev
}
