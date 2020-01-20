package mystack

import (
	"testing"

	"Leetcode/algorithms/kit"

	"github.com/stretchr/testify/assert"
)

// run: go test -v base.go 173.*
func Test_BSTIterator(t *testing.T) {
	assert := assert.New(t)

	ints := [][]int{
		{7, 3, 15, NULL, NULL, 9, 20},
		{7, 3, 15, 1, 2, 9, 20},
	}
	for _, arr := range ints {
		root := kit.Ints2Tree(arr)
		nums := kit.Tree2Inorder(root)
		t.Logf("%v trans to tree and inorder is: %v\n", arr, nums)

		iterator := NewBSTIterator(root)
		for i := 0; iterator.HasNext(); i++ {
			assert.Equal(nums[i], iterator.Next(), "i: %d", i)
		}
	}

}
