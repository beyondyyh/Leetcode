package tree

// TreeNode declare
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL declare
var NULL = -1 << 63

// breadthFirstSearch 层级遍历
func breadthFirstSearch(root TreeNode) []interface{} { // {{{
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
} // }}}

/**
 * 构造tree
 *	when step = 0:
 *					1
 *				2			5
 *			3		4	6		7
 */
func buildTree(step int) *TreeNode { // {{{
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
} // }}}

// Ints2Tree 利用[]int生成*TreeNode
func Ints2Tree(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: ints[0]}
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	for i := 1; i < n; {
		node := queue[0]
		queue = queue[1:]
		if ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		// 注意i可能越界
		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}
