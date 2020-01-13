package main

import "math"

// run: go test -v 7.*
func reverse(x int) int {
	var res int
	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}
	if res >= math.MaxInt32 || res <= math.MinInt32 {
		return 0
	}
	return res
}
