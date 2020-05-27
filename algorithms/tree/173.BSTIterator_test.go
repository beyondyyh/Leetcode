package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"Leetcode/algorithms/kit"
)

// run: go test -v -run Test_BSTIterator
func Test_BSTIterator(t *testing.T) {
	assert := assert.New(t)
	ints := []int{7, 3, 15, null, null, 9, 20}
	root := kit.Ints2Tree(ints)

	nums := kit.Tree2Inorder(root)

	var i int
	it := NewBSTIterator(root)
	for it.HasNext() {
		assert.Equal(nums[i], it.Next(), "%d", nums[i])
		i++
	}
}
