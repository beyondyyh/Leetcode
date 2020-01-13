package mystring

import (
	"fmt"
	"path/filepath"
	"strings"

	"gopl.io/interview2020/Leetcode/algorithms/kit"
)

// 1. 用栈
// 2. 先把字符串以"/"为分隔符分割成数组，此数组包含"路径字符串"、""、"."、".."这四种情况;
// 3. 遍历数组, 当s[i]==".."并且栈不空时pop, 当s[i]!="" && s[i]!="." && s[i]!=".."时，把s[i]路径压入栈;
// 4. 如果栈空, 返回"/", 栈非空, 从栈底部（注意不是栈顶）开始拼接字符串
func simplifyPath(path string) string {
	stack := kit.NewStack()
	parts := strings.Split(path, "/")
	fmt.Printf("origin:%s split parts:%v\n", path, parts)

	for _, part := range parts {
		if !stack.IsEmpty() && part == ".." {
			stack.Pop()
		} else if part != "" && part != "." && part != ".." {
			stack.Push(part)
		}
	}

	if stack.IsEmpty() {
		return "/"
	}

	var items []string
	for i := 0; i < stack.Len(); i++ {
		items = append(items, (stack.Seek(i)).(string))
	}
	return "/" + strings.Join(items, "/")
}

// golang标准库提供了牛逼的此方法
func cleanFilepath(path string) string {
	return filepath.Clean(path)
}
