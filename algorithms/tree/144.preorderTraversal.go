package tree

import "gopl.io/interview2020/Leetcode/algorithms/kit"

// preorderTraversal 基于栈的前序遍历
// 前序遍历顺序：根->左->右，在遍历左、右子树时，仍然先访问根结点，然后遍历左子树，最后遍历右子树
// 步骤：访问根结点 -> 前序遍历左子树 -> 前序遍历右子树，遍历左右子树时仍然采用前序遍历方法
// 如以下二叉树:
//				A
//		B				C
// D		E		F		NULL
// 前序遍历：ABDECF		中序遍历：DBEAFC	后序遍历：DEBFCA
// 时间复杂度：O(n), 空间复杂度：O(n)
func preorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	stack := kit.NewStack()
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
