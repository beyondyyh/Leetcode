package mystack

// Stack 存放interface{}的栈通用定义
type Stack struct {
	elements []interface{}
}

// NewStack 返回*mystack.Stack
func NewStack() *Stack {
	// []interface{}初始化需要make
	// []int{}不需要
	return &Stack{elements: make([]interface{}, 0)}
}

// Len 返回栈的长度
func (s *Stack) Len() int {
	return len(s.elements)
}

// Push 将x放入栈
func (s *Stack) Push(x interface{}) {
	s.elements = append(s.elements, x)
}

// Pop 弹出栈顶部 element
func (s *Stack) Pop() interface{} {
	x := s.elements[s.Len()-1]
	s.elements = s.elements[:s.Len()-1]
	return x
}

// Peek 查看栈顶部 element，并不修改内部结构
func (s *Stack) Peek() interface{} {
	x := s.elements[s.Len()-1]
	return x
}

// Seek 从栈底部依次取出 element, 不修改内部结构
func (s *Stack) Seek(i int) interface{} {
	if i < 0 {
		i = 0
	}
	if i > s.Len()-1 {
		i = s.Len() - 1
	}
	return s.elements[i]
}

// IsEmpty 返回s是否为空
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
