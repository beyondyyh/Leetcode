package myqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v 622.*
func Test_MyCircularQueue(t *testing.T) {
	assert := assert.New(t)
	q := NewMyCircularQueue(3)
	t.Logf("the Q is: %+v\n", q)

	// Q is empty
	assert.True(q.IsEmpty(), "Check Q is empty")
	assert.False(q.IsFull(), "Check Q is full")
	assert.False(q.DeQueue(), "Check DeQueue when Q is empty")
	assert.Equal(-1, q.Front(), "Check the front when Q is empty")
	assert.Equal(-1, q.Rear(), "Check the rear when Q is empty")

	// Insert
	start, end := 1, 3
	for i := start; i <= end; i++ {
		assert.True(q.EnQueue(i), "Check insert %d into Q", i)
	}
	// Q is full
	assert.False(q.EnQueue(4), "Check insert into Q when is full")
	assert.False(q.IsEmpty(), "Check Q is empty")
	assert.True(q.IsFull(), "Check Q is full")
	assert.Equal(start, q.Front(), "Check the front of Q")
	assert.Equal(end, q.Rear(), "Check the rear of Q")

	assert.True(q.DeQueue(), "Check delete from Q")
	assert.Equal(2, q.Front(), "Check the front of Q")
	assert.Equal(3, q.Rear(), "Check the rear of Q")
	t.Logf("the Q is: %+v\n", q)

	q.DeQueue()
	t.Logf("the Q is: %+v\n", q)
	assert.Equal(3, q.Front(), "Check the front of Q")
	assert.Equal(3, q.Rear(), "Check the rear of Q")

	q.DeQueue()
	t.Logf("the Q is: %+v\n", q)
	assert.Equal(-1, q.Front(), "Check the front of Q")
	assert.Equal(-1, q.Rear(), "Check the rear of Q")
	assert.True(q.IsEmpty(), "Check Q is empty")
	assert.False(q.IsFull(), "Check Q is full")
}
