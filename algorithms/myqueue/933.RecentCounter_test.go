package myqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v 933.*
func Test_RecentCounter(t *testing.T) {
	assert := assert.New(t)
	rc := NewRecentCounter()

	times := []int{1, 100, 3001, 3002}
	counters := []int{1, 2, 3, 3}

	for i := range times {
		t, c := times[i], counters[i]
		assert.Equal(c, rc.Ping(t))
	}
}
