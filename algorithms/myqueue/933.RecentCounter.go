package myqueue

import (
	"Leetcode/algorithms/kit"
)

// 题目：最近的请求次数 https://leetcode-cn.com/problems/number-of-recent-calls/
// 题解：此题主要的理解题目吧，理解了就很简单
// 只会考虑最近3000毫秒到现在的 ping 数，因此可以使用队列存储这些 ping 的记录。
// 当收到一个时间 t 的 ping 时，我们将它加入队列，并且将所有在时间 t - 3000 之前的 ping 移出队列，
// 同时返回队列的长度即可

// RecentCounter structure declare
type RecentCounter struct {
	queue *kit.Queue
}

// NewRecentCounter
func NewRecentCounter() *RecentCounter {
	return &RecentCounter{
		queue: kit.NewQueue(),
	}
}

// Ping returns ping times within the last 3000
func (rc *RecentCounter) Ping(t int) int {
	rc.queue.Push(t)

	// 把3000毫秒以外ping的 t 移除队列
	for rc.queue.Peek() < t-3000 {
		rc.queue.Pop()
	}

	return rc.queue.Len()
}
