package mystack

import "gopl.io/interview2020/Leetcode/algorithms/kit"

// https://leetcode-cn.com/problems/binary-search-tree-iterator/
//		7
// 3		15
//		9		20
// 二叉搜索树的前序遍历

// TreeNode is pre-defined

// BSTIterator binary-search-tree iterator structure declare
type BSTIterator struct {
	stack *Stack
}

// NewBSTIterator returns *BSTIterator
func NewBSTIterator(root *TreeNode) *BSTIterator {
	it := &BSTIterator{
		stack: kit.NewStack(),
	}
	it.mypush(root)
	return it
}

// Next 取出栈顶元素node。把node的右子树的左节点全部入栈
func (it *BSTIterator) Next() int {
	node := (it.stack.Pop()).(*TreeNode)
	if node.Right != nil {
		it.mypush(node.Right)
	}
	return node.Val
}

// HasNext 只要栈不空就是还有下一个元素
func (it *BSTIterator) HasNext() bool {
	return !it.stack.IsEmpty()
}

// mypush 遍历当前节点的左子树，一直找到leaf节点
func (it *BSTIterator) mypush(node *TreeNode) {
	for node != nil {
		it.stack.Push(node)
		node = node.Left
	}
}
