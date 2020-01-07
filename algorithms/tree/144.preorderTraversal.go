package tree

import (
	"gopl.io/interview2020/Leetcode/algorithms/mystack"
)

// preorderTravelsal 基于栈的前序遍历 顺序：根->左->右
// Top->Bottom 和 Left->Right， 入栈时先入右子树
// 时间复杂度：O(n), 空间复杂度：O(n)
func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	stack := mystack.NewStack()
	stack.Push(root)
	for !stack.IsEmpty() {
		// curr node
		node := (stack.Pop()).(*TreeNode)
		res = append(res, node.Val)

		// 栈先入后出，访问时是先访问左子树->右子树，
		// 故此处先push右子树
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}

	}
	return res
}

// 递归大法
func preorderRecurse(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	res = append(res, root.Val)
	res = append(res, preorderRecurse(root.Left)...)
	res = append(res, preorderRecurse(root.Right)...)
	return res
}
