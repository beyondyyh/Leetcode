package tree

import (
	"testing"
)

// run: go test -v treeNode.go 617.*
func Test_mergeTrees(t *testing.T) {
	t1 := buildTree(0)
	t2 := buildTree(100)
	mt := mergeTrees(t1, t2)
	t.Logf("t1:%v t2:%v mt:%v \n", breadthFirstSearch(*t1), breadthFirstSearch(*t2), breadthFirstSearch(*mt))
	// output: t1:[102 104 110 106 108 112 114] t2:[101 102 105 103 104 106 107] mt:[102 104 110 106 108 112 114]
}
