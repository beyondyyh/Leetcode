package mystack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v 225.*
func Test_MyStack(t *testing.T) {
	assert := assert.New(t)

	s := NewMyStack()
	assert.True(s.Empty(), "MyStack should be empty")

	start, end := 1, 10
	for i := start; i <= end; i++ {
		s.Push(i)
		assert.Equal(i, s.Top(), "Check the top element after push")
	}

	for i := end; i >= start; i-- {
		x := s.Pop()
		t.Logf("%d ", x)
		assert.Equal(i, x, "Check pop from s")
	}
	assert.True(s.Empty(), "MyStack should be empty")
}
