package myheap

import (
	"container/heap"
)

// 第K大的数，用最小堆
// 时间复杂度：O(N(logK))
// 空间复杂度：O(K)
func findKthLargest(nums []int, k int) int {
	if k <= 0 || len(nums) == 0 {
		return 0
	}

	h := &MinHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}
