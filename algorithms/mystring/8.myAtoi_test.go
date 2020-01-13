package main

import (
	"math"
	"testing"
)

// run: go test -v 8.*
func Test_myAtoi(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name: "正数",
			input: "	  42 ", // 1 \t and 2 space with leading and 1 space with tailing
			want: 42,
		},
		{
			name:  "负数",
			input: "    -42",
			want:  -42,
		},
		{
			name:  "数字开头",
			input: "4193 with words",
			want:  4193,
		},
		{
			name:  "字母开头",
			input: "words and 987",
			want:  0,
		},
		{
			name: "无效数字",
			input: "	 ",
			want: 0,
		},
		{
			name:  "超过32位有符号整数范围",
			input: "-91283472332",
			want:  math.MinInt32,
		},
		{
			name:  "超过32位无符号整数范围",
			input: "91283472332",
			want:  math.MaxInt32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.input); got != tt.want {
				t.Errorf("myAtoi(%v)=%v, want=%v", tt.input, got, tt.want)
			}
		})
	}
}
