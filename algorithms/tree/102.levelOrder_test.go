package tree

import (
	_ "fmt"
	"testing"
)

// [3,9,20,nil,nil,15,7]
func build_levelOrder_case() *TreeNode {
	var root *TreeNode = new(TreeNode)
	root.Val = 3
	var l1 *TreeNode = new(TreeNode)
	l1.Val = 9

	var r1 *TreeNode = new(TreeNode)
	r1.Val = 20
	var r2l *TreeNode = new(TreeNode)
	r2l.Val = 15
	var r2r *TreeNode = new(TreeNode)
	r2r.Val = 7
	r1.Left = r2l
	r1.Right = r2r

	root.Left = l1
	root.Right = r1
	return root
}

func Test_BFS(t *testing.T) {
	tests := build_levelOrder_case()
	t.Logf("case: %v\nBFS: %v\n levelOrder: %v\n",
		breadthFirstSearch(*tests), BFS(tests), levelOrder(tests))
}
