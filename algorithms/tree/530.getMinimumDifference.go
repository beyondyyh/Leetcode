package tree

import (
	"Leetcode/algorithms/kit"
)

func getMinimumDifference(root *TreeNode) int {
	// DFS: 中序遍历
	ints := kit.Tree2Inorder(root)

	var res = 0
	for i := 1; i < len(ints); i++ {
		if res == 0 || ints[i]-ints[i-1] < res {
			res = ints[i] - ints[i-1]
		}
	}
	return res
}
