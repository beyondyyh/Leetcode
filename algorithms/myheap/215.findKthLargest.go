package myheap

import (
	"container/heap"
	"fmt"
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
		fmt.Printf("%v push\n", num)
		if h.Len() > k {
			x := heap.Pop(h)
			fmt.Printf("%v pop\n", x)
		}
		fmt.Printf("heap length is:%d\n", h.Len())
	}
	return (*h)[0]
}
