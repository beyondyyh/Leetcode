package tree

import (
	"testing"

	"gopl.io/interview2020/Leetcode/algorithms/kit"
)

// run: go test -v base.go 617.*
func Test_mergeTrees(t *testing.T) {
	t0 := kit.Ints2Tree([]int{1, 2, 3, 4, 5, 6, 7})
	t1 := t0
	t2 := kit.Ints2Tree([]int{100, 101, 102, 103, 104, 105, 106, 107})
	mt := mergeTrees(t0, t2)
	t.Logf("t1:%v t2:%v mt:%v \n", kit.BreadthFirstSearch(*t1), kit.BreadthFirstSearch(*t2), kit.BreadthFirstSearch(*mt))
	// outout: t1:[101 103 105 107 109 111 113 107] t2:[100 101 102 103 104 105 106 107] mt:[101 103 105 107 109 111 113 107]
}
