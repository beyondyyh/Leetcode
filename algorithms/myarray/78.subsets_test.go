package myarray

import (
	"reflect"
	"sort"
	"testing"
)

// run: go test -v 78.*
func Test_subsets(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:     "[]",
			input:    []int{},
			expected: [][]int{nil},
		},
		{
			name:     "[1]",
			input:    []int{1},
			expected: [][]int{nil, []int{1}},
		},
		{
			name:     "[1,2,3]",
			input:    []int{1, 2, 3},
			expected: [][]int{nil, []int{1}, []int{2}, []int{3}, []int{1, 2}, []int{1, 3}, []int{2, 3}, []int{1, 2, 3}},
		},
		{
			name:  "[9,0,3,5]",
			input: []int{9, 0, 3, 5},
			expected: [][]int{nil, []int{9}, []int{9, 0}, []int{9, 0, 3}, []int{9, 3}, []int{9, 0, 3, 5}, []int{9, 5}, []int{9, 0, 5},
				[]int{9, 3, 5}, []int{0}, []int{0, 3}, []int{0, 3, 5}, []int{0, 5}, []int{3}, []int{3, 5}, []int{5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			/*
				if output := subsets(tt.input); !reflect.DeepEqual(output, tt.expected) {
					t.Errorf("subsets(%v)=%v, expected=%v", tt.input, output, tt.expected)
				}
			*/

			output := subsets(tt.input)
			expected := tt.expected
			sortSlice(output)
			sortSlice(expected)
			// t.Logf("subsets(%v)=%v, expected=%v", tt.input, output, tt.expected)
			if !reflect.DeepEqual(output, expected) {
				t.Errorf("subsets(%v)=%v, expected=%v", tt.input, output, tt.expected)
			}
		})
	}
}

func sortSlice(iss [][]int) {
	// 每个二维数组排序
	for i, _ := range iss {
		sort.Ints(iss[i])
	}

	sort.SliceStable(iss, func(i, j int) bool {
		if len(iss[i]) == len(iss[j]) {
			for k := range iss[i] {
				if iss[i][k] == iss[j][k] {
					continue
				}
				return iss[i][k] < iss[j][k]
			}
		}
		return len(iss[i]) < len(iss[j])
	})
}
