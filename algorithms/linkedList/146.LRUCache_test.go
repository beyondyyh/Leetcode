package linkedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v base.go 146.*
// capacity is 3
func Test_LRUCache_c3(t *testing.T) { // {{{
	assert := assert.New(t)

	cache := Constructor(3)
	cache.Put(0, 0)
	cache.Put(1, 1)
	cache.Put(2, 2)
	// [(2,2), (1,1), (0,0)]

	assert.Equal(1, cache.Get(1), "get 1 from c3")
	// [(1,1), (2,2), (0,0)]

	cache.Put(3, 3)
	// [(3,3), (1,1), (2,2)]

	assert.Equal(-1, cache.Get(0), "get 0 from c3")
	// 未变, [(3,3), (1,1), (2,2)]

	cache.Put(4, 4)
	// [(4,4), (3,3), (1,1)]
	assert.Equal(-1, cache.Get(2), "get 2 from c3")
	assert.Equal(3, cache.Get(3), "get 3 from c3")
	// [(3,3), (4,4), (1,1)]

	assert.Equal(4, cache.Get(4), "get 4 from c3")
	// [(4,4), (3,3), (1,1)]

	cache.Put(1, 100)                               // [(1,100), (4,4), (3,3)]
	cache.Put(5, 5)                                 // [(5,5), (1,100), (4,4)]
	assert.Equal(-1, cache.Get(3), "get 3 from c3") // 3 not exists

	assert.Equal(100, cache.Get(1), "get 1 from c3")
	// [(1,100), (5,5), (4,4)]
} // }}}

// capacity is 1
func Test_LRUCache_c1(t *testing.T) { // {{{
	assert := assert.New(t)

	cache := Constructor(1)
	cache.Put(0, 0)
	// [(0,0)]
	cache.Put(1, 1)
	// [(1,1)]
	cache.Put(2, 2)
	// [(2,2)]
	assert.Equal(-1, cache.Get(1), "get 1 from c1")
	cache.Put(3, 3)
	// [(3,3)]
	cache.Put(3, 33)
	// [(3,33)]
	assert.Equal(33, cache.Get(3), "get 3 from c1")
} // }}}
