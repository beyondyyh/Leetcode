package tree

func zigzagLevelOrder2(root *TreeNode) [][]int {
	var res [][]int
	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		// 出现新的 level
		if level >= len(res) {
			res = append(res, []int{})
		}
		// 与 102 题相比，level 的奇偶不同，append 会不同
		if level%2 == 0 {
			res[level] = append(res[level], root.Val)
		} else {
			tmp := make([]int, len(res[level])+1)
			tmp[0] = root.Val
			copy(tmp[1:], res[level])
			res[level] = tmp
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return res
}
