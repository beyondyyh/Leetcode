package kit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v queue*
func Test_Queue(t *testing.T) {
	assert := assert.New(t)

	q := NewQueue()
	assert.True(q.IsEmpty(), "Check q is empty or not")

	start, end := 1, 10
	for i := start; i <= end; i++ {
		q.Push(i)
		assert.Equal(i-start+1, q.Len(), "Check q.Len() after push")
	}

	for i := start; i <= end; i++ {
		x := q.Pop()
		t.Logf("q.Pop() = %d\n", x)
		assert.Equal(i, x, "Pop from q")
	}
	assert.True(q.IsEmpty(), "Check q is empty or not after pop")
}
