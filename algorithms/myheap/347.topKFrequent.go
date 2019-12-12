package myheap

import (
	"container/heap"
	"fmt"
)

// 堆排序，实现Len(), Less(), Swap(), Push(), Pop()方法
// topK问题，用最小堆
type entry struct {
	Count int // 出现次数
	Value int // 元素值
}

type MaxHeapTopK []entry

func (h MaxHeapTopK) Len() int {
	return len(h)
}

func (h MaxHeapTopK) Less(i, j int) bool {
	return h[i].Count < h[j].Count
}

func (h MaxHeapTopK) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeapTopK) Push(x interface{}) {
	*h = append(*h, x.(entry))
}

func (h *MaxHeapTopK) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 解题思路：
// 借助 哈希表 来建立数字和其出现次数的映射，遍历一遍数组统计元素的频率
// 维护一个元素数目为 kk 的最小堆
// 每次都将新的元素与堆顶元素（堆中频率最小的元素）进行比较
// 如果新的元素的频率比堆顶端的元素大，则弹出堆顶端的元素，将新的元素添加进堆中
// 最终，堆中的 kk 个元素即为前 kk 个高频元素
func topKFrequent(nums []int, k int) []int {
	if k <= 0 || len(nums) == 0 {
		return nil
	}

	// 利用map计数
	var dict = make(map[int]int)
	for _, num := range nums {
		if _, ok := dict[num]; ok {
			dict[num]++
		} else {
			dict[num] = 1
		}
	}

	fmt.Printf("dict is: %v\n", dict)

	// push入堆，当堆长度＞k时，将堆顶元素弹出
	h := &MaxHeapTopK{}
	heap.Init(h)
	for val, cnt := range dict {
		entry := entry{Count: cnt, Value: val}
		heap.Push(h, entry)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	fmt.Printf("heap's length is: %d\n", h.Len())

	// 遍历输出
	var res []int
	for h.Len() > 0 {
		e := heap.Pop(h).(entry)
		fmt.Println(e)
		// res = append(res, heap.Pop(h).(entry).Value)
		res = append(res, e.Value)
	}
	return res
}
