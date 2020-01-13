package linkedList

import (
	"testing"

	"gopl.io/interview2020/Leetcode/algorithms/kit"
)

// run: go test -v base.go 234.*
func Test_isPalindrome(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "偶数回文",
			input:    []int{1, 2, 3, 2, 1},
			expected: true,
		},
		{
			name:     "偶数回文",
			input:    []int{1, 2, 2, 1},
			expected: true,
		},
		{
			name:     "非回文",
			input:    []int{1, 2, 3, 2},
			expected: false,
		},
		{
			name:     "单节点回文",
			input:    []int{1},
			expected: true,
		},
		{
			name:     "空节点回文",
			input:    []int{},
			expected: true,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			head := kit.Ints2List(tt.input)
			if output := isPalindrome(head); output != tt.expected {
				t.Errorf("isPalindrome(%v)=%v, expected=%v", tt.input, output, tt.expected)
			}
		})
	}
}
