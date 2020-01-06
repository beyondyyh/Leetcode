package main

import (
	"fmt"
	"math"

	"gopl.io/interview2020/Leetcode/algorithms/mystack"
)

// Stack redeclare
// type Stack mystack.Stack

// 方法一：暴力法
// 使用栈，遍历原始字符串找到每种可能的非空偶数长度的字符串，然后利用栈逐个判断该字串是否是合法的，非常暴力！
// 时间复杂度：O(n²)，空间复杂度：O(n)
func longestValidParentheses1(s string) int {
	var maxLen float64 = 0
	for i := 0; i < len(s); i++ {
		for j := i + 2; j <= len(s); {
			substr := string([]byte(s)[i:j])
			fmt.Printf("substr:%s, isValid:%t\n", substr, isValid(substr))
			if isValid(substr) {
				// math.Max, math.Min参数和返回值都是float64类型
				maxLen = math.Max(maxLen, float64(j-1))
			}
			j = j + 2
		}
	}
	return int(maxLen)
}

// 此处默认s只包含 '(' 和 ')'
func isValid(s string) bool {
	stack := mystack.NewStack()
	for _, c := range []byte(s) {
		if c == '(' {
			// 如果当前ele是'('则压入栈
			stack.Push('(')
		} else if !stack.IsEmpty() && stack.Peek() == '(' {
			// 如果当前ele是')'，则看下栈顶是否是'('，因为只有'('才能与之配对，是则弹出
			// 需要注意Pop之前先判断栈是否为空，否则会数组越界
			stack.Pop()
		} else {
			return false
		}
	}
	// 最后判断栈是否为空，为空则说明全部匹配完成，isValid
	return stack.IsEmpty()
}
