package mystack

import (
	_ "fmt"

	"Leetcode/algorithms/kit"
)

// MyStack Declare
type MyStack struct {
	queue *Queue
}

// NewMyStack Constructor
func NewMyStack() *MyStack {
	return &MyStack{
		queue: kit.NewQueue(),
	}
}

// 思路：先将目标值入队列，再将目标值前面的n-1个数输出再依次压进入队尾
// 时间复杂度：O(n) 空间复杂度：O(1)
// Push element x onto stack
func (s *MyStack) Push(x int) {
	// fmt.Printf("before push:%v\n", s.queue.PeekAll())
	s.queue.Push(x)
	// 从0到n-1元素
	for i := 0; i < s.queue.Len()-1; i++ {
		top := s.queue.Peek() // 队头元素
		s.queue.Pop()
		s.queue.Push(top)
	}
	// fmt.Printf("after push:%v\n", s.queue.PeekAll())
}

// Pop Removes the element on top of the stack and returns that element
func (s *MyStack) Pop() int {
	if s.queue.IsEmpty() {
		return NULL
	}

	top := s.queue.Peek()
	s.queue.Pop()
	return top
}

// Top Get the top element
func (s *MyStack) Top() int {
	if s.queue.IsEmpty() {
		return NULL
	}

	return s.queue.Peek()
}

// Empty Returns whether the stack is empty
func (s *MyStack) Empty() bool {
	return s.queue.IsEmpty()
}
