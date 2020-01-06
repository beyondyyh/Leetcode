package main

import (
	"fmt"
	"strings"

	"gopl.io/interview2020/Leetcode/algorithms/mystack"
)

// run: go test -v 71.*
func simplifyPath(path string) string {
	stack := mystack.NewStack()
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
