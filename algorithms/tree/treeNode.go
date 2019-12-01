package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

/**
 * 构造tree
 *	when step = 0:
 *					1
 *				2			5
 *			3		4	6		7
 */
func buildTree(step int) *TreeNode {
	var root *TreeNode = new(TreeNode)
	root.Val = step + 1
	var l1 *TreeNode = new(TreeNode)
	l1.Val = step + 2
	var l2 *TreeNode = new(TreeNode)
	l2.Val = step + 3
	var l2r *TreeNode = new(TreeNode)
	l2r.Val = step + 4
	l1.Left = l2
	l1.Right = l2r

	var r1 *TreeNode = new(TreeNode)
	r1.Val = step + 5
	var r2l *TreeNode = new(TreeNode)
	r2l.Val = step + 6
	var r2 *TreeNode = new(TreeNode)
	r2.Val = step + 7
	r1.Left = r2l
	r1.Right = r2

	root.Left = l1
	root.Right = r1
	return root
}
