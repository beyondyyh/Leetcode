package kit

// TreeNode declare
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL declare
var NULL = -1 << 63

// BreadthFirstSearch 层级遍历
func BreadthFirstSearch(root TreeNode) []interface{} { // {{{
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

// Ints2Tree 利用[]int生成*TreeNode
// ints 和广度优先遍历BFS结果相同
func Ints2Tree(ints []int) *TreeNode { // {{{
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
} // }}}

// Tree2Inorder 把 二叉树转换成 inorder 的切片
// 递归大法 顺序：左->根->右
func Tree2Inorder(root *TreeNode) []int { // {{{
	var res []int
	if root == nil {
		return res
	}

	res = Tree2Inorder(root.Left)
	res = append(res, root.Val)
	res = append(res, Tree2Inorder(root.Right)...)
	return res
} // }}}

// Tree2Preorder 把 二叉树转换成 preorder 的切片
// 递归大法 顺序：根->左->右
func Tree2Preorder(root *TreeNode) []int { // {{{
	var res []int
	if root == nil {
		return res
	}

	res = append(res, root.Val)
	res = append(res, Tree2Preorder(root.Left)...)
	res = append(res, Tree2Preorder(root.Right)...)
	return res
} // }}}

// Tree2Postorder 把 二叉树转换成 postorder 的切片
// 递归大法 顺序：左->右->根
func Tree2Postorder(root *TreeNode) []int { // {{{
	var res []int
	if root == nil {
		return res
	}

	res = append(res, Tree2Postorder(root.Left)...)
	res = append(res, Tree2Postorder(root.Right)...)
	res = append(res, root.Val)
	return res
} // }}}
