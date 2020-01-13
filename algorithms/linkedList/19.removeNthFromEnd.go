package linkedList

// 方法一：简化为删除第L-n+1个节点，两次遍历法
// 删除从列表开头数起的第 (L - n + 1) 个结点，其中 L 是列表的长度
// 时间复杂度：O(L)，对列表进行了两次遍历，首先计算了列表的长度 L 其次找到第 (L-n) 个结点。操作执行了 2L-n 步，时间复杂度为 O(L)
// 空间复杂度：O(1)
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	// 申明一个哑结点，该结点位于列表头部，处理极端情况，例如列表中只含有一个结点，或需要删除列表的头部
	var dummy *ListNode = &ListNode{Val: 0}
	dummy.Next = head
	curr := head
	var L int = 0
	for curr != nil {
		L++
		curr = curr.Next
	}

	// 第L-n个节点
	L -= n
	curr = dummy // 重置到头结点
	for L > 0 {
		curr = curr.Next
		L--
	}
	curr.Next = curr.Next.Next
	return dummy.Next
}

// 方法二：双指针法，只遍历一次
// 申明两个指针fast、slow同时指向头部，fast指针先走n步，然后一起移动，此时二者的距离为n，
// 当fast到尾部nil时，slow的位置恰好为倒数第n个节点
// 因为要删除该节点，所以要移动到该节点的前一个才能删除，所以循环结束条件为 fast.next != null
// 时间复杂度：O(n)
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	var dummy *ListNode = &ListNode{Val: 0}
	dummy.Next = head
	fast, slow := dummy, dummy

	// fast先走n步
	for n > 0 {
		fast = fast.Next
		n--
	}

	// if n > len(head), remove the head
	if fast == nil {
		return dummy.Next.Next
	}

	// fast、slow同时开始移动
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// 此时slow所处位置正好是倒数第n个节点的前一个
	slow.Next = slow.Next.Next
	return dummy.Next
}
