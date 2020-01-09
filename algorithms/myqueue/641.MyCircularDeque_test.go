package myqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v 641.*
func Test_MyCircularDeque(t *testing.T) {
	assert := assert.New(t)
	q := NewMyCircularDeque(3)
	t.Logf("q is: %+v\n", q) // &{front:0 rear:0 length:4 nums:[0 0 0 0]}

	// Q is empty
	assert.False(q.IsFull(), "检查空队列是否满")
	assert.True(q.IsEmpty(), "检查空队列是否空")
	assert.False(q.DeleteFront(), "从空队列中删除首部元素")
	assert.False(q.DeleteLast(), "从空队列中删除尾部元素")
	assert.Equal(-1, q.GetFront(), "从空队列中获取首部元素")
	assert.Equal(-1, q.GetRear(), "从空队列中获取尾部元素")

	// when i=10 InsertFront之后：[10, 10，10]
	for i := 1; i <= 10; i++ {
		assert.True(q.InsertFront(i), "在首部插入 %d", i)
		assert.True(q.InsertFront(i), "在首部插入 %d", i)
		assert.True(q.InsertFront(i), "在首部插入 %d", i)
		assert.True(q.DeleteLast(), "删除尾部元素 %d", i)
		assert.True(q.DeleteLast(), "删除尾部元素 %d", i)
		assert.True(q.DeleteLast(), "删除尾部元素 %d", i)
	}
	t.Logf("q is: %+v\n", q) // &{front:2 rear:2 length:4 nums:[10 9 10 10]}
	assert.True(q.IsEmpty(), "检查空队列是否空")

	for i := 100; i <= 110; i++ {
		assert.True(q.InsertLast(i), "在尾部插入 %d", i)
		assert.True(q.InsertLast(i), "在尾部插入 %d", i)
		assert.True(q.InsertLast(i), "在尾部插入 %d", i)
		assert.True(q.DeleteFront(), "删除首部元素 %d", i)
		assert.True(q.DeleteFront(), "删除首部元素 %d", i)
		assert.True(q.DeleteFront(), "删除首部元素 %d", i)
	}
	t.Logf("q is: %+v\n", q) // &{front:3 rear:3 length:4 nums:[110 110 110 109]}
	assert.True(q.IsEmpty(), "检查空队列是否空")

	// 初始化下
	q = NewMyCircularDeque(3)
	t.Logf("q is: %+v\n", q) // &{front:0 rear:0 length:4 nums:[0 0 0 0]}
	assert.True(q.InsertLast(1), "往尾部插入 1")
	assert.True(q.InsertLast(2), "往尾部插入 2")
	assert.True(q.InsertFront(3), "往首部插入 3")
	t.Logf("q is: %+v\n", q) // &{front:3 rear:2 length:4 nums:[1 2 0 3]}

	assert.True(q.IsFull(), "检查满队列")
	assert.False(q.IsEmpty(), "检查满队列")

	assert.False(q.InsertFront(4), "往首部插入 4")
	assert.False(q.InsertLast(4), "往尾部插入 4")
	t.Logf("q is: %+v\n", q) // &{front:3 rear:2 length:4 nums:[1 2 0 3]}

	assert.Equal(2, q.GetRear(), "获取尾部元素")
	assert.Equal(3, q.GetFront(), "获取首部元素")

	assert.True(q.DeleteLast(), "删除尾部元素")
	t.Logf("q is: %+v\n", q) // &{front:3 rear:1 length:4 nums:[1 2 0 3]}
	assert.True(q.InsertFront(4), "在首部添加 4")
	t.Logf("q is: %+v\n", q) // &{front:2 rear:1 length:4 nums:[1 2 4 3]}

	assert.Equal(4, q.GetFront(), "检查首部元素")
}
