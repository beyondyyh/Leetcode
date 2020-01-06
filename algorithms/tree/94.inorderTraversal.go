package tree

import (
	"gopl.io/interview2020/Leetcode/algorithms/mystack"
)

// 基于栈的遍历
// 时间复杂度：O(n), 空间复杂度：O(n)
func inorderTraversal(root *TreeNode) []int {
	var res []int
	stack := mystack.NewStack()
	curr := root

	for curr != nil || !stack.IsEmpty() {
		for curr != nil {
			stack.Push(curr)
			curr = curr.Left
		}
		curr = (stack.Pop()).(*TreeNode)
		res = append(res, curr.Val)
		curr = curr.Right
	}

	return res
}

// 递归遍历
func inorderRecurse(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := inorderRecurse(root.Left)
	res = append(res, root.Val)
	res = append(res, inorderRecurse(root.Right)...)

	return res
}
