package tree

type TreeNode struct {
	Data  interface{}
	Left  *TreeNode
	Right *TreeNode
}

// 广度优先遍历
func BreadthFirstSearch(root TreeNode) []interface{} {
	var res []interface{}
	var nodes []TreeNode = []TreeNode{root}
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		res = append(res, node.Data)
		if node.Left != nil {
			nodes = append(nodes, *node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, *node.Right)
		}
	}
	return res
}
