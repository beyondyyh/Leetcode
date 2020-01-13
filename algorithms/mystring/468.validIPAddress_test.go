package mystring

import "testing"

// run: go test -v 468.*
func Test_validIPAddress(t *testing.T) {
	cases := []struct {
		name, input, expected string
	}{
		{"x1", "172.16.254.1", "IPv4"},
		{"x2", "2001:0db8:85a3:0:0:8A2E:0370:7334", "IPv6"},
		{"x3", "256.256.256.256", "Neither"},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if output := validIPAddress(tt.input); output != tt.expected {
				t.Errorf("validIPAddress(%s)=%s, expected=%s", tt.input, output, tt.expected)
			}
		})
	}
}
