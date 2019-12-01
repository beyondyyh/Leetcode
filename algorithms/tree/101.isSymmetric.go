package tree

func isSymmetric(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}

	if t1.Val == t2.Val {
		return isSymmetric(t1.Left, t2.Right) && isSymmetric(t1.Right, t2.Left)
	}
	return false
}
