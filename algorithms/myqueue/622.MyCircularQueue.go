package myqueue

// 题目：设计循环队列 https://leetcode-cn.com/problems/design-circular-queue/
// 设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以形成一个循环。它也被称为“环形缓冲器”。
// 循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入下一个元素，即使在队列前面仍有空间。但是使用循环队列，我们能使用这些空间去存储新的值。
//
// 你的实现应该支持如下操作：
//	MyCircularQueue(k): 构造器，设置队列长度为 k 。
// 		Front: 从队首获取元素。如果队列为空，返回 -1 。
// 		Rear: 获取队尾元素。如果队列为空，返回 -1 。
// 		enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
// 		deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
// 		isEmpty(): 检查循环队列是否为空。
// 		isFull(): 检查循环队列是否已满。

// 题解：参考liweiwei1419非常棒的设计思路（https://leetcode-cn.com/problems/design-circular-queue/solution/shu-zu-shi-xian-de-xun-huan-dui-lie-by-liweiwei141/），基于golang实现

import "fmt"

// MyCircularQueue structure declare
type MyCircularQueue struct {
	front  int   // 指向队列头部第1个有效数据的位置
	rear   int   // 指向队列尾部（即最后1个有效数据）的下一个位置，即下一个从队尾入队元素的位置
	length int   // 队列长度，非容量，slice动态库容
	nums   []int // 队列元素
}

// NewMyCircularQueue Initialize the structure. Set the size of the queue to be k.
func NewMyCircularQueue(k int) *MyCircularQueue {
	length := k + 1
	return &MyCircularQueue{
		length: length,
		nums:   make([]int, length),
	}
}

// Front Returns the first item of the queue
func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}

	return q.nums[q.front]
}

// Rear Returns the last item of the queue.
func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}

	// 其实就是rear-1
	pos := (q.rear - 1 + q.length) % q.length
	return q.nums[pos]
}

// EnQueue Insert an element into the queue.
// Returns true if the operation is successful.
func (q *MyCircularQueue) EnQueue(x int) bool {
	fmt.Printf("front:%d rear:%d IsFull:%t\n", q.front, q.rear, q.IsFull())
	if q.IsFull() {
		return false
	}

	q.nums[q.rear] = x
	q.rear = (q.rear + 1) % q.length // rear指针后移一位
	return true
}

// DeQueue Delete an element from the queue.
// Returns true if the operation is successful.
func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	q.front = (q.front + 1) % q.length // front指针后移一位
	return true
}

// IsEmpty check the queue is empty.
func (q *MyCircularQueue) IsEmpty() bool {
	return q.front == q.rear
}

// IsFull check the queue is full.
func (q *MyCircularQueue) IsFull() bool {
	// 这是这个经典设计的原因
	return (q.rear+1)%q.length == q.front
}
