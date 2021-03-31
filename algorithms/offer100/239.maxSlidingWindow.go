package offer100

// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
// 返回滑动窗口中的最大值。
// 示例 1：
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
// 解释：
// 滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7
// Leetcode ：https://leetcode-cn.com/problems/sliding-window-maximum

// 可以采用最大堆的数据结构来保存元素，堆顶元素即为当前堆的最大值，并判断当前堆顶元素这是否在窗口中，
// 在则直接返回，不在则利用堆特性删除堆顶元素并调整堆
// 时间复杂度：O(N logN)， n=len(nums)

import (
	"container/heap"
	"fmt"
)

type Item struct {
	index int // 元素下标
	value int // 元素值
}

// 声名一个优先队列，java的优先队列就是用堆实现的
type PriorityQueue []*Item

// 实现heap接口的 Len,Less,Swap,Push,Pop 方法
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// 比较器，当两者的值相同时，比较下标的位置
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].value == pq[j].value {
		return pq[i].index > pq[j].index
	}
	return pq[i].value > pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push、Pop不仅改变堆的值还改变堆长度，所以receiver用指针
func (pq *PriorityQueue) Push(x interface{}) {
	// n := len(*pq)
	item := x.(*Item)
	// item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // 避免内存泄露
	*pq = old[0 : n-1]
	return item
}

// 只查看堆顶元素，不更改堆的结构，此处是大顶堆所以第一个元素最大
func (pq PriorityQueue) peek() interface{} {
	return pq[0]
}

func dump(pq PriorityQueue) {
	nums := make([]Item, 0)
	for _, v := range pq {
		nums = append(nums, *v)
	}
	fmt.Printf("pq: %+v\n", nums)
}

// -----
func maxSlidingWindow(nums []int, k int) []int {
	// 初始化前K的元素到堆中
	pq := make(PriorityQueue, k)
	for i := 0; i < k; i++ {
		pq[i] = &Item{index: i, value: nums[i]}
	}
	heap.Init(&pq)
	// dump(pq)

	n := len(nums)
	// 总共有n-k+1个窗口，声名一个长度长度&容量n-k+1的slice
	ans := make([]int, n-k+1)
	// 堆顶元素即是第一个窗口最大值，先放进ans
	ans[0] = pq.peek().(*Item).value

	// 遍历将剩下的元素依次入堆
	for i := k; i < n; i++ {
		// 将新元素入堆
		item := &Item{index: i, value: nums[i]}
		heap.Push(&pq, item)
		// 循环判断当前堆顶是否在窗口中，一般思路是遍历窗口元素与堆顶进行对比，时间复杂度为O(k)
		// 反向思维：堆顶元素已经是最大值，可以依次pop比较堆顶元素的下标是否小于窗口的左边界i-k+1，直到堆为空或者堆顶元素下标等于左边界，出栈时间复杂度O(1)
		for pq.Len() > 0 && pq.peek().(*Item).index <= i-k {
			heap.Pop(&pq)
		}
		// 在窗口中直接赋值即可
		if pq.Len() > 0 {
			ans[i-k+1] = pq.peek().(*Item).value
		}
	}
	// dump(pq)
	return ans
}
