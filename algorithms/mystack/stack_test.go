package mystack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v stack*
func Test_Stack(t *testing.T) {
	assert := assert.New(t)

	s := NewStack()
	assert.True(s.IsEmpty(), "Check s is empty or not")

	start, end := 1, 10
	for i := start; i <= end; i++ {
		s.Push(i)
		assert.Equal(i-start+1, s.Len(), "Check s.Len() after push")
	}
	for i := end; i >= start; i-- {
		x := s.Pop()
		t.Logf("%d ", x)
		assert.Equal(i, x, "Pop from s")
	}
	assert.True(s.IsEmpty(), "Check s is empty or not after pop")

	t.Log("********")

	for i := 0; i < 10; i++ {
		letter := string(i + 'A')
		s.Push(letter)
		assert.Equal(i+1, s.Len(), "Check s.Len() after push")
	}
	for i := 9; i >= 0; i-- {
		x := s.Pop()
		t.Logf("%s ", x)
		assert.Equal(string(i+'A'), x, "Pop from s")
	}
}
