package main

import "fmt"

func main() {
	s := "abcdefg"
	stack := []byte(s)
	fmt.Println(pop(byte('g'), &stack), stack) // true
	fmt.Println("---------")
	fmt.Println(isValid("()()((([]{})))")) // true
	fmt.Println(isValid("()()())"))        // false
}

func isValid(s string) bool {
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
}

// 实现一个栈 的pop方法 栈：先入后出
func pop(char byte, stack *[]byte) bool {
	l := len(*stack)
	if l < 1 {
		return false
	}

	// 注意stack是引用类型。。
	var poped byte
	*stack, poped = (*stack)[:l-1], (*stack)[l-1]
	return poped == char
}
