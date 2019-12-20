package mystack

// Stack 存放int数组的栈
type Stack struct {
	nums []int
}

// NewStack 返回*mystack.Stack
func NewStack() *Stack {
	return &Stack{nums: []int{}}
}

// Len 返回s的长度
func (s *Stack) Len() int {
	return len(s.nums)
}

// Push 将n放入栈
func (s *Stack) Push(n int) {
	s.nums = append(s.nums, n)
}

// Pop 弹出栈顶element
func (s *Stack) Pop() int {
	x := s.nums[s.Len()-1]
	s.nums = s.nums[:s.Len()-1]
	return x
}

// IsEmpty 返回s是否为空
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}
