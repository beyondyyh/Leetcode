package linkedList

// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例：
//
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	var head *ListNode = &ListNode{Val: 0}
	cur := head

	var carry int
	for l1 != nil || l2 != nil {
		// 踩坑：变量作用域的问题，xy一定不能再for外面申明，否则，
		// 如果len(*l2) < len(*l1)时，y值异常，本来应该是0，但是保留了上一个位置的数
		var x, y, sum int
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		sum = x + y + carry
		carry = sum / 10
		sum = sum % 10

		// cur 循环
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	// 最后一位
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return head.Next
}
