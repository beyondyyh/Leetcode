package tree

// 二叉搜索树（Binary Search Tree）又叫 `二叉排序树` 他是一颗空树 或者具有以下特点：
// 1. 针对某个节点，根节点上的值大于等于左子树的值，小于等于右子树的值
// 2. 他的左右子树也分别是二叉排序树
// 通过 `中序遍历` 可以得到一个升序的数组，这点很重要

import (
	"Leetcode/algorithms/kit"
)

type BSTIterator struct {
	index int
	nums  []int
}

func NewBSTIterator(root *TreeNode) *BSTIterator {
	return &BSTIterator{
		index: -1,
		nums:  kit.Tree2Inorder(root),
	}
}

func (it *BSTIterator) Next() int {
	it.index += 1
	return it.nums[it.index]
}

func (it *BSTIterator) HasNext() bool {
	return it.index+1 < len(it.nums)
}
