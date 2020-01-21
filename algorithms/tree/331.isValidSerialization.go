package tree

import (
	"strings"

	"Leetcode/algorithms/kit"
)

// 读取到的结构只要符合二叉树性质而且不会在未读完之前就满足leaves = nodes + 1（完整的二叉树）即可
func isValidSerialization1(preorder string) bool {
	var leaves, node int
	pres := strings.Split(preorder, ",")
	for i, s := range pres {
		if s == "#" {
			leaves++
		} else {
			node++
		}
		if leaves > node+1 {
			return false
		}
		if leaves == node+1 && i < len(pres)-1 {
			return false
		}
	}

	if leaves == node+1 {
		return true
	}

	return false
}

// 基于stack
// 使用栈来模拟二叉树的遍历，就是如果此时栈中的前两层都是null（"#"），那就将这颗子树弹出去，
// 过程中要判断这棵子树的根节点是否为空，为空直接返回false，如果不为空将这棵树弹出去之后用一个"#"代替，
// 这样一遍遍历下来栈中保留的就应该只有一个"#",对其进行判断，满足返回true，否则返回false
func isValidSerialization2(preorder string) bool {
	if len(preorder) == 0 {
		return false
	}

	array := strings.Split(preorder, ",")
	stack := kit.NewStack()
	for _, ele := range array {
		stack.Push(ele)
		for stack.Len() >= 3 && stack.Peek().(string) == "#" && stack.Seek(stack.Len()-2).(string) == "#" {
			stack.Pop()
			stack.Pop()
			// 如果节点为空，返回false
			if stack.Peek().(string) == "#" {
				return false
			}
			stack.Pop()

			stack.Push("#")
		}
	}

	return stack.Len() == 1 && stack.Peek().(string) == "#"
}
