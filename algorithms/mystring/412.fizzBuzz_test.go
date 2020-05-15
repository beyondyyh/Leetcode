package mystring

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"Leetcode/algorithms/kit"
)

// go test -v -run Test_fizzBuzz
func Test_fizzBuzz(t *testing.T) {
	cases := []kit.CaseEntry{
		{
			Name:     "x1",
			Input:    1,
			Expected: []string{"1"},
		},
		{
			Name:     "x2",
			Input:    5,
			Expected: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			Name:     "x3",
			Input:    15,
			Expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.Name, func(t *testing.T) {
			output := fizzBuzz(tt.Input.(int))
			expected := tt.Expected.([]string)
			assert.Equal(t, expected, output, "equal")
		})
	}
}
