package linkedList

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gopl.io/interview2020/Leetcode/algorithms/kit"
)

// run: go test -v base.go 148.*
func Test_sortList(t *testing.T) {
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
		t.Logf("~~%v~~\n", tt)
		head := kit.Ints2List(tt.input)
		assert.Equal(tt.expected, kit.List2Ints(sortList(head)), "输入:%v", tt)
	}
}

// run: go test -bench=Benchmark_sortList -benchmem -count=3
// 1. 参数-bench，要测试的函数；点字符测试当前所有以Benchmark为前缀函数
// 2. 参数-benchmem，性能测试时显示测试函数的内存分配大小，内存分配次数的统计信息
// 3. 参数-count n，运行测试和性能多少此，默认1
func Benchmark_sortList(b *testing.B) {
	head := kit.Ints2List([]int{9, -1, 3, 1, 4, 8, 2, 5, 7, 19, 16, 13, 11, 14, 18, 12, 15, 17, 29, 26, 23, 21, 24, 28, 22, 0, 100, 999})
	b.Logf("~~~~~%v~~~~~", b.N)
	for i := 0; i < b.N; i++ {
		sortList(head)
	}
}
