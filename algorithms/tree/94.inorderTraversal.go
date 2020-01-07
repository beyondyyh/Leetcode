package tree

import (
	"gopl.io/interview2020/Leetcode/algorithms/mystack"
)

// inorderTraversal 基于栈的遍历 顺序：左->根->右
// 时间复杂度：O(n), 空间复杂度：O(n)
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	stack := mystack.NewStack()
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

// 递归大法 顺序：左->根->右
func inorderRecurse(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	res = inorderRecurse(root.Left)
	res = append(res, root.Val)
	res = append(res, inorderRecurse(root.Right)...)
	return res
}
