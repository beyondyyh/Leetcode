package tree

// 判断一个树是否存在链表的路径
func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}

	// 先判断当前的节点，如果不对，再看左右子树
	return isMatch(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

// 判断二叉树中的每个节点为起点往下的路径是否与链表相匹配
// head表示当前已经匹配到的链表节点，root表示二叉树的节点
func isMatch(head *ListNode, root *TreeNode) bool {
	// 特判1：链表走完了，返回true
	if head == nil {
		return true
	}

	// 特判2：链表没走完，树走完了，这肯定不行啊，返回false
	if root == nil {
		return false
	}

	// 特判3：值不同，必定不是啊
	if head.Val != root.Val {
		return false
	}

	// 正常逻辑：如果值相同，继续走，左子树和右子树有一个满足即可
	return isMatch(head.Next, root.Left) || isMatch(head.Next, root.Right)
}
