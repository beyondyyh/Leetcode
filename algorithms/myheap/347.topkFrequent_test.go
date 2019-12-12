package myheap

import "testing"

func Test_topkFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3, 4, 4, 4, 4, 4, 4}
	k := 2
	res := topKFrequent(nums, k)
	t.Logf("nums is: %v\n topk is: %v\n", nums, res)
}
