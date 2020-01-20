package binarysearch

import (
	"testing"

	"Leetcode/algorithms/kit"
)

// run: go test -v 69.*
func Test_mySqrt(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:     "x1",
			Input:    1,
			Expected: 1,
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			if output := mySqrt(tt.Input.(int)); output != tt.Expected.(int) {
				t.Errorf("mySqrt(%d)=%d, expected=%d", tt.Input, output, tt.Expected)
			}
		})
	}
}
