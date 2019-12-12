package myheap

import (
	"container/heap"
	_ "fmt"
)

func maxHeapDemo() (max interface{}, res []interface{}) {
	h := &MaxHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	max = (*h)[0]
	for h.Len() > 0 {
		res = append(res, heap.Pop(h))
	}
	return
}

func minHeapDemo() (min interface{}, res []interface{}) {
	h := &MinHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	min = (*h)[0]
	for h.Len() > 0 {
		res = append(res, heap.Pop(h))
	}
	return
}
