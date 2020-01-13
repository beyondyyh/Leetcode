package mystring

import (
	"reflect"
	"sort"
	"testing"
)

// run: go test -v 93.*
func Test_restoreIpAddress(t *testing.T) {
	cases := []struct {
		name, input string
		expected    []string
	}{
		{
			name:     "x1",
			input:    "25525511135",
			expected: []string{"255.255.11.135", "255.255.111.35"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			output := restoreIpAddresses(tt.input)
			sort.Strings(output)
			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("restoreIpAddresses(%s)=%v, expected=%v", tt.input, output, tt.expected)
			}
		})
	}
}
