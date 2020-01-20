package tree

import (
	"testing"

	"Leetcode/algorithms/kit"
)

// run: go test -v 331.*
func Test_isValidSerialization(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:     "x1",
			Input:    "9,3,4,#,#,1,#,#,2,#,6,#,#",
			Expected: true,
		},
		{
			Name:     "x2",
			Input:    "1,#",
			Expected: false,
		},
		{
			Name:     "x3",
			Input:    "9,#,#,1",
			Expected: false,
		},
		// more test case
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			if output := isValidSerialization(tt.Input.(string)); output != tt.Expected.(bool) {
				t.Errorf("isValidSerialization(%s)=%t, expected=%t", tt.Input, output, tt.Expected.(bool))
			}
		})
	}
}
