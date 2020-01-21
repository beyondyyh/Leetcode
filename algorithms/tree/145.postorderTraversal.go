package tree

import (
	"Leetcode/algorithms/kit"
)

// postorderTraversal 基于栈的后序遍历
// 后序遍历顺序：左->右->根，在遍历左、右子树时，仍然先遍历左子树，然后遍历右子树，最后遍历根结点
// 步骤：后序遍历左子树 -> 后序遍历右子树 -> 根节点，遍历左右子树时仍然采用后序遍历方法
// 如以下二叉树:
//				A
//		B				C
// D		E		F		NULL
// 前序遍历：ABDECF		中序遍历：DBEAFC	后序遍历：DEBFCA
// 时间复杂度：O(n), 空间复杂度：O(n)
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	// 保存最近访问的一个的根节点，用于判断当前是从左子树到的根节点还是右子树到的根节点
	// 如果当前节点的右节点和上一次遍历的节点相同，那就表明当前是从右节点过来的了
	var latest *TreeNode = nil
	stack := kit.NewStack()

	for root != nil || !stack.IsEmpty() {
		if root != nil {
			// 与中序遍历一样，都需要先找到左子树的leaf节点
			stack.Push(root)
			root = root.Left
		} else {
			// curr node
			curr := (stack.Peek()).(*TreeNode)
			// fmt.Printf("curr:%v latest:%v\n", curr, latest)
			// 是否变到右子树
			if curr.Right != nil && curr.Right != latest {
				root = curr.Right
			} else {
				res = append(res, curr.Val)
				latest = curr
				stack.Pop()
			}
		}
	}
	return res
}
