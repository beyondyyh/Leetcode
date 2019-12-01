package tree

import (
	"fmt"
	"testing"
)

// buildx returns [1 2 2 3 4 4 3]
func buildx() *TreeNode {
	head := new(TreeNode)
	head.Val = 1
	// left node
	l1 := new(TreeNode)
	l1.Val = 2
	l1.Left = new(TreeNode)
	l1.Left.Val = 3
	l1.Right = new(TreeNode)
	l1.Right.Val = 4
	// right node
	r1 := new(TreeNode)
	r1.Val = 2
	r1.Left = new(TreeNode)
	r1.Left.Val = 4
	r1.Right = new(TreeNode)
	r1.Right.Val = 3

	head.Left = l1
	head.Right = r1
	return head
}

func Test_isSymmetric(t *testing.T) {
	tt := buildx()
	t.Logf("t: %v\n", breadthFirstSearch(*tt))
	fmt.Println(isSymmetric(tt, tt))
}
