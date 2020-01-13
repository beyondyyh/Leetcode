package kit

// 最大堆
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

// 最大堆，h[i] > h[j]
// 最小堆，h[i] < h[j]
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push and Pop use pointer receivers because they modify the slice's length,
// not just its contents.
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
