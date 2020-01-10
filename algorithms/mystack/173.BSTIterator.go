package mystack

//
// import (
// 	_ "gopl.io/interview2020/Leetcode/algorithms/tree"
// )
//
// // TreeNode is pre-defined
// // var TreeNode = tree.TreeNode
//
// // BSTIterator binary-search-tree iterator structure declare
// type BSTIterator struct {
// 	stack *Stack
// }
//
// func NewBSTIterator(root *TreeNode) *BSTIterator {
// 	it := &BSTIterator{
// 		stack: NewStack(),
// 	}
// 	it.stack.Push(root)
// 	return it
// }
//
// func (it *BSTIterator) Next() int {
// 	node := it.stack.Pop()
// 	if node.Right != nil {
// 		it.push(node.Right)
// 	}
// 	return node.Val
// }
//
// func (it *BSTIterator) HasNext() bool {
// 	return !it.stack.IsEmpty()
// }
//
// func (it *BSTIterator) push(node *TreeNode) {
// 	for node != nil {
// 		it.stack.Push(node)
// 		node = node.Left
// 	}
// }
