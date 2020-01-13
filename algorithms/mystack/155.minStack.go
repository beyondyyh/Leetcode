package mystack

import "gopl.io/interview2020/Leetcode/algorithms/kit"

// MinStack declare
type MinStack struct {
	data   *Stack // 数据栈
	helper *Stack // 辅助栈
}

// NewMinStack returns MinStack constructor
func NewMinStack() *MinStack {
	return &MinStack{
		data:   kit.NewStack(),
		helper: kit.NewStack(),
	}
}

// key1: 辅助栈的元素空的时候，必须放入新进来的数
// key2: 新来的数小于或者等于辅助栈栈顶元素的时候，才放入
// key3: 出栈时，辅助栈的栈顶元素等于数据栈的栈顶元素，才出栈，即"出栈保持同步"
func (s *MinStack) Push(x int) {
	s.data.Push(x)
	if s.helper.IsEmpty() || (s.helper.Peek()).(int) >= x {
		s.helper.Push(x)
	}
}

// Pop pop the top ele of stack
func (s *MinStack) Pop() {
	if !s.data.IsEmpty() {
		top := s.data.Pop()
		if top == s.helper.Peek() {
			s.helper.Pop()
		}
	}
}

// Top returns the top ele of stack, not modify internal struct
func (s *MinStack) Top() int {
	if !s.data.IsEmpty() {
		return (s.data.Peek()).(int)
	}
	panic("Illegal operation Top(), stack is empty")
}

// GetMin returns the minimum ele of stack
func (s *MinStack) GetMin() int {
	if !s.helper.IsEmpty() {
		return (s.helper.Peek()).(int)
	}
	panic("Illegal operation GetMin(), stack is empty")
}
