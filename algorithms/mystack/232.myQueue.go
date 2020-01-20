package mystack

import "Leetcode/algorithms/kit"

// problem: https://leetcode-cn.com/problems/implement-queue-using-stacks/
// 使用栈实现队列的下列操作：
//
// push(x) -- 将一个元素放入队列的尾部。
// pop() -- 从队列首部移除元素。
// peek() -- 返回队列首部的元素。
// empty() -- 返回队列是否为空。

// 题解：
// 1、使用两个栈，（stackPush）用于元素进栈，（stackPop）用于元素出栈；
// 2、pop() 或 peek() 的时候：
//	（1）如果 stackPop 里面有元素，直接从 stackPop 里弹出或者 peek 元素；
//	（2）如果 stackPop 里面没有元素，一次性将 stackPush 里面的所有元素倒入 stackPop

// MyQueue declare
type MyQueue struct {
	stackPush, stackPop *Stack
}

// NewMyQueue constructor
func NewMyQueue() *MyQueue {
	return &MyQueue{
		stackPush: kit.NewStack(),
		stackPop:  kit.NewStack(),
	}
}

// Push Push element x to the back of queue
func (q *MyQueue) Push(x int) {
	q.stackPush.Push(x)
}

// Pop Removes the element from in front of queue and returns that element.
func (q *MyQueue) Pop() int {
	q.shift()
	if !q.stackPop.IsEmpty() {
		return (q.stackPop.Pop()).(int)
	}
	panic("Illegal operation Pop(), queue is empty")
}

// Peek Returns the front element
func (q *MyQueue) Peek() int {
	q.shift()
	if !q.stackPop.IsEmpty() {
		return (q.stackPop.Peek()).(int)
	}
	panic("Illegal operation Peek(), queue is empty")
}

// Empty Returns whether the queue is empty
func (q *MyQueue) Empty() bool {
	return q.stackPush.IsEmpty() && q.stackPop.IsEmpty()
}

// shift 一次性将 stackPush 里的所有元素倒入stackPop中
//	1、该操作只在 stackPop 里为空的时候才执行，否则会破坏出队入队的顺序
//	2、在 peek 和 pop 操作之前调用该方法
func (q *MyQueue) shift() {
	if !q.stackPop.IsEmpty() {
		return
	}

	for !q.stackPush.IsEmpty() {
		q.stackPop.Push(q.stackPush.Pop())
	}
}
