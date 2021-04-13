package myarray

import (
	"fmt"
	"sort"
)

// 方法一：最容易理解，先把nums2 append到nums1尾部，然后利用快排思想排序
// 时间复杂度：快排平均时间复杂度 O(m+n)log(m+n)
// 空间复杂度：快排平均空间复杂度 O(m+n)log(m+n)
func merge1(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2) // func copy(dst, src []Type) int
	sort.Ints(nums1)
	fmt.Println(nums1)
}

// 方法二：双指针法，每次从两个数组头部取出比较小的数字放到结果中，需要引入额外数组
// 时间复杂度：O(m+n)
// 空间复杂度：O(m+n) 引入了额外的空间
func merge2(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			sorted = append(sorted, nums1[i])
			i++
		} else {
			sorted = append(sorted, nums2[j])
			j++
		}
	}
	// 未遍历完的情况
	if i < m {
		sorted = append(sorted, nums1[:i]...)
	}
	if j < n {
		sorted = append(sorted, nums2[:j]...)
	}
	copy(nums1, sorted)
}

// 方法二更优雅的写法
func merge3(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	i, j := 0, 0
	for {
		if i == m { // nums1遍历结束，将nums2剩下的项append到sorted
			sorted = append(sorted, nums2[j:]...)
			break
		}
		if j == n { // nums2遍历结束，将nums1剩下的项append到sorted
			sorted = append(sorted, nums1[i:]...)
			break
		}

		if nums1[i] <= nums2[j] {
			sorted = append(sorted, nums1[i])
			i++
		} else {
			sorted = append(sorted, nums2[j])
			j++
		}
	}
	copy(nums1, sorted)
}

// 方法三：方法二的优化版，逆向双指针
// 方法二之所以用一个merged临时数组来保存，是因为如果直接合并nums2到nums1中则会覆盖还没遍历到的部分。
// 优化思路：可以将nums2先放到nums1后面空的部分这样直接覆盖不会影响结果，因此可以指针从后往前遍历，每次取2者之中较大的放进nums1的最后面。
// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func merge4(nums1 []int, m int, nums2 []int, n int) {
	i, j, tail := m-1, n-1, m+n-1 // i:nums1的尾部，j:nums2的尾部，tail:合并的尾部
	for j >= 0 {
		if i < 0 || nums1[i] <= nums2[j] {
			nums1[tail] = nums2[j]
			j--
		} else {
			nums1[tail] = nums1[i]
			i--
		}
		tail--
	}
}
