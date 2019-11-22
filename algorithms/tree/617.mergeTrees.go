package tree

// 617. 合并二叉树
func mergeTrees(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

// breadthFirstSearch 层级遍历
func breadthFirstSearch(root TreeNode) []interface{} {
	var res []interface{}
	var nodes []TreeNode = []TreeNode{root}
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			nodes = append(nodes, *node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, *node.Right)
		}
	}
	return res
}
