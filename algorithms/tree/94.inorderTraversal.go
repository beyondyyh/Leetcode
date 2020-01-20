package tree

import (
	"Leetcode/algorithms/kit"
)

// inorderTraversal 基于栈的中序遍历
// 中序遍历顺序：左->根->右
// 步骤：中序遍历左子树 -> 访问根结点 -> 中序遍历右子树
// 如以下二叉树:
//				A
//		B				C
// D		E		F		NULL
// 前序遍历：ABDECF		中序遍历：DBEAFC	后序遍历：DEBFCA
// 时间复杂度：O(n), 空间复杂度：O(n)
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	stack := kit.NewStack()
	for root != nil || !stack.IsEmpty() {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		// curr node
		node := (stack.Pop()).(*TreeNode)
		res = append(res, node.Val)
		root = node.Right
	}

	return res
}
