package tree

import (
	"testing"
)

/**
 * 构造结果
 *	when step = 0:
 *					1
 *				2			5
 *			3		4	6		7
 */
func buildTree(step int) *TreeNode {
	var root *TreeNode = new(TreeNode)
	root.Val = step
	var l1 *TreeNode = new(TreeNode)
	l1.Val = step + 1
	var l2 *TreeNode = new(TreeNode)
	l2.Val = step + 2
	var l2r *TreeNode = new(TreeNode)
	l2r.Val = step + 3
	l1.Left = l2
	l1.Right = l2r

	var r1 *TreeNode = new(TreeNode)
	r1.Val = step + 4
	var r2l *TreeNode = new(TreeNode)
	r2l.Val = step + 5
	var r2 *TreeNode = new(TreeNode)
	r2.Val = step + 6
	r1.Left = r2l
	r1.Right = r2

	root.Left = l1
	root.Right = r1
	return root
}

// run: go test -v treeNode.go 617.*
func Test_mergeTrees(t *testing.T) {
	t1 := buildTree(1)
	t2 := buildTree(101)
	mt := mergeTrees(t1, t2)
	t.Logf("t1:%v t2:%v mt:%v \n", breadthFirstSearch(*t1), breadthFirstSearch(*t2), breadthFirstSearch(*mt))
	// output: t1:[102 104 110 106 108 112 114] t2:[101 102 105 103 104 106 107] mt:[102 104 110 106 108 112 114]
}
