package mystring

import "Leetcode/algorithms/kit"

// isValid1 based on stack
func isValid1(s string) bool { // {{{
	stack := make([]byte, 0)
	smap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range []byte(s) {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if v, ok := smap[char]; ok {
				// 出栈第一个元素，if出栈失败，返回false
				// 注意此处用的引用
				if !pop(v, &stack) {
					return false
				}
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
} // }}}

// 实现一个栈 的pop方法 栈：先入后出
func pop(char byte, stack *[]byte) bool { // {{{
	l := len(*stack)
	if l < 1 {
		return false
	}

	// 注意stack是引用类型。。
	var poped byte
	*stack, poped = (*stack)[:l-1], (*stack)[l-1]
	return poped == char
} // }}}

// isValid2 based on kit.Stack
func isValid2(s string) bool {
	stack := kit.NewStack()
	smap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range []byte(s) {
		if c == '(' || c == '[' || c == '{' {
			stack.Push(c)
		} else {
			v, ok := smap[c]
			// invalid char
			if !ok {
				return false
			}
			if stack.IsEmpty() {
				return false
			}

			// top of the stacke is not in ['(', '[', '{']
			if (stack.Peek()).(byte) != v {
				return false
			}
			stack.Pop()
		}
	}

	return stack.IsEmpty()
}
