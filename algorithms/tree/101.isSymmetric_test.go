package tree

import (
	"fmt"
	"testing"

	"gopl.io/interview2020/Leetcode/algorithms/kit"
)

// go test -v base.go 101.*
func Test_isSymmetric(t *testing.T) {
	tt := kit.Ints2Tree([]int{1, 2, 2, 3, 4, 4, 3})
	t.Logf("t: %v\n", kit.BreadthFirstSearch(*tt))
	fmt.Println(isSymmetric(tt, tt))
}
