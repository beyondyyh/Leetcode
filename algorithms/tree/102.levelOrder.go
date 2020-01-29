package tree

// BFS 普通的BFS遍历
func BFS(root *TreeNode) []int { // {{{
	var res []int
	var queue []*TreeNode = []*TreeNode{root}
	if root == nil {
		return res
	}

	for len(queue) > 0 {
		node := queue[0]
		res = append(res, node.Val)
		queue = queue[1:]

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
} // }}}

// 1. 普通的BFS 比较简单 只需要什么一个[]int 一个存放tree的队列 利用queue先进先出的规则打印value即可
// 2. 如果根据等级来保存[][]int的话 只要声明一个正在遍历的队列 一个保存下一层节点的队列即可
func levelOrder1(root *TreeNode) [][]int { // {{{
	var res [][]int
	if root == nil {
		return res
	}

	var queue []*TreeNode = []*TreeNode{root}
	for len(queue) > 0 {
		var tmpQueue []*TreeNode
		var tmpRes []int
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			tmpRes = append(tmpRes, node.Val)
			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}
			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
		}
		queue = queue[0:0]
		queue = append(queue, tmpQueue...)
		res = append(res, tmpRes)
	}

	return res
} // }}}

// 递归大法
func levelOrder2(root *TreeNode) [][]int {
	var res [][]int
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		// 出现了新的level
		if level >= len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)
	return res
}
