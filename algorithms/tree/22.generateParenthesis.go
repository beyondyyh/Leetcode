package tree

import _ "fmt"

func generateParenthesis(n int) []string {
	var res []string
	if n == 0 {
		return res
	}

	// dfs隐式回溯
	dfs("", 0, 0, n, &res)
	return res
}

/**
 * 深度优先遍历法
 * @ref https://leetcode-cn.com/problems/generate-parentheses/solution/hui-su-suan-fa-by-liweiwei1419/
 * @param curstr 当前递归得到的结果
 * @param left   左括号用了几个
 * @param right  右括号用了几个
 * @param n      左括号、右括号一共用几个
 * @param res    结果集
 */
func dfs(curstr string, left, right, n int, res *[]string) {
	// fmt.Printf("left:%d right:%d n:%d curstr:%s\n", left, right, n, curstr)
	// 因为是递归函数，所以先写递归终止条件
	if left == n && right == n {
		*res = append(*res, curstr)
		// fmt.Printf("curstr:%s res:%v\n", curstr, *res)
		return
	}

	// 如果左括号还没凑够，继续凑
	if left < n {
		// 拼接上一个左括号，并且剩余的左括号个数减 1
		dfs(curstr+"(", left+1, right, n, res)
	}

	// 什么时候可以用右边？例如，(((((()，此时 left > right，
	// 不能用等号，因为只有先拼了左括号，才能拼上右括号
	if right < n && left > right {
		dfs(curstr+")", left, right+1, n, res)
	}
}
