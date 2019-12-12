package myheap

import "testing"

// run: go test -v minHeap.go maxHeap.go 215.*
func Test_findKthLargest(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	x := findKthLargest(nums, k)
	t.Logf("%v the %dth largest is: %d\n", nums, k, x)
}
