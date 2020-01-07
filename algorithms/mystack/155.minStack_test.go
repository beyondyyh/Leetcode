package mystack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v stack.go 155.*
func Test_MinStack(t *testing.T) {
	assert := assert.New(t)

	s := NewMinStack()
	s.Push(-2)
	s.Push(0)
	s.Push(-3) // [-2, 0, -3]
	assert.Equal(s.GetMin(), -3, "The min should be -3")

	s.Pop() // [-2, 0]
	assert.Equal(s.Top(), 0, "The top should be 0")

	s.Push(1) // [-2, 0, 1]
	assert.Equal(s.GetMin(), -2, "The min should be -2")

	s.Pop() // [-2, 0]
	s.Pop() // [-2]
	s.Pop() // []
	assert.Panics(func() {
		s.Top()
	}, "The Top() shoud be panic.")
	assert.PanicsWithValue("Illegal operation GetMin(), stack is empty", func() {
		s.GetMin()
	}, "The GetMin() should be panic.")
}
