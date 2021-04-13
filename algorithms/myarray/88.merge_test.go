package myarray

import (
	"reflect"
	"testing"
)

type (
	entry88input struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	entry88 struct {
		name     string
		input    entry88input
		expected []int
	}
)

var cases = []entry88{
	{
		name: "x1",
		input: entry88input{
			nums1: []int{1, 2, 3, 0, 0, 0}, // len=3, cap=6
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
		},
		expected: []int{1, 2, 2, 3, 5, 6},
	},
	{
		name: "x2",
		input: entry88input{
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
		},
		expected: []int{1},
	},
}

// run: go test -run Test_merge1
func Test_merge1(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			merge1(c.input.nums1, c.input.m, c.input.nums2, c.input.n)
			if !reflect.DeepEqual(c.input.nums1, c.expected) {
				t.Errorf("merge1(%v, %d, %v, %d)=%v, expected=%v", c.input.nums1, c.input.m, c.input.nums2, c.input.n, c.input.nums1, c.expected)
			}
		})
	}
}

// run: go test -run Test_merge2
func Test_merge2(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			merge2(c.input.nums1, c.input.m, c.input.nums2, c.input.n)
			if !reflect.DeepEqual(c.input.nums1, c.expected) {
				t.Errorf("merge2(%v, %d, %v, %d)=%v, expected=%v", c.input.nums1, c.input.m, c.input.nums2, c.input.n, c.input.nums1, c.expected)
			}
		})
	}
}

// run: go test -run Test_merge3
func Test_merge3(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			merge3(c.input.nums1, c.input.m, c.input.nums2, c.input.n)
			if !reflect.DeepEqual(c.input.nums1, c.expected) {
				t.Errorf("merge3(%v, %d, %v, %d)=%v, expected=%v", c.input.nums1, c.input.m, c.input.nums2, c.input.n, c.input.nums1, c.expected)
			}
		})
	}
}
