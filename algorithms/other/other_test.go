package other

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// run: go test -v -run Test_quickSort
func Test_quickSort(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{4, 2, 1, 3},
			[]int{1, 2, 3, 4},
		},
		{
			[]int{-1, 5, 3, 4, 0},
			[]int{-1, 0, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		output := quickSort(tt.input)
		assert.Equal(tt.expected, output)
	}
}

// run: go test -v -run Test_mergeSort
func Test_mergeSort(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{4, 2, 1, 3},
			[]int{1, 2, 3, 4},
		},
		{
			[]int{-1, 5, 3, 4, 0},
			[]int{-1, 0, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		output := mergeSort(tt.input)
		assert.Equal(tt.expected, output)
	}
}
