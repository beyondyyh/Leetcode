package mystring

import (
	"testing"

	"Leetcode/algorithms/kit"
)

// go test -v -run Test_titleToNumber
func Test_titleToNumber(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:     "x1",
			Input:    "A",
			Expected: 1,
		},
		{
			Name:     "x2",
			Input:    "AB",
			Expected: 28,
		},
		{
			Name:     "x3",
			Input:    "ZY",
			Expected: 701,
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			if output := titleToNumber(tt.Input.(string)); output != tt.Expected.(int) {
				t.Errorf("titleToNumber(%v)=%d, expected=%v", tt.Input, output, tt.Expected)
			}
		})
	}
}
