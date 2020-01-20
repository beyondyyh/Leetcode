package tree

import (
	_ "fmt"
	"testing"

	"Leetcode/algorithms/kit"
)

// go test -v base.go 102.*
func Test_levelOrder(t *testing.T) {
	ints := []int{3, 9, 20, NULL, NULL, 15, 7}
	tests := kit.Ints2Tree(ints)
	t.Logf("case: %v\nBFS: %v\n levelOrder: %v\n",
		kit.BreadthFirstSearch(*tests), BFS(tests), levelOrder(tests))
}
