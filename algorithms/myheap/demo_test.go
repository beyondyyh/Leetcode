package myheap

import (
	"testing"
)

func Test_maxHeapDemo(t *testing.T) {
	max, res := maxHeapDemo()
	t.Logf("max:%v res:%v\n", max, res)
}

func Test_minHeapDemo(t *testing.T) {
	min, res := minHeapDemo()
	t.Logf("min:%v res:%v\n", min, res)
}
