package mystack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v base.go 232.*
func Test_MyQueue(t *testing.T) {
	assert := assert.New(t)

	q := NewMyQueue()
	assert.True(q.Empty(), "Check q is empty")

	start, end := 1, 10
	for i := start; i <= end; i++ {
		q.Push(i)
	}
	assert.Equal(start, q.Peek(), "Check the front element of q")

	for i := start; i <= end; i++ {
		x := q.Pop()
		t.Logf("q.Pop() = %d\n", x)
		assert.Equal(i, x, "Check pop from q")
	}
	assert.True(q.Empty(), "Check q is empty after all pop")

	// q now is: []
	assert.Panics(func() {
		q.Peek()
	}, "The Peek() shoud be panic.")

	assert.PanicsWithValue("Illegal operation Pop(), queue is empty", func() {
		q.Pop()
	}, "The Pop() should be panic.")
}
