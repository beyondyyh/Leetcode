package other

// 归并排序是利用归并的思想，采用分治策略
// 分治：分是将大问题分成一些小问题然后递归求解，而治则是将分的阶段得到的答案“合并”在一起
func mergeSort(nums []int) []int {
	n := len(nums)
	if n <= 1 {
		return nums // 最后切割只剩下一个元素
	}

	// 合并2个有序数组，双指针
	var _merge func(left, right []int) []int
	_merge = func(left, right []int) []int {
		m, n := len(left), len(right)
		res := make([]int, 0, m+n)

		var i, j int
		for i < m && j < n {
			if left[i] < right[j] {
				res = append(res, left[i])
				i++
			} else {
				res = append(res, right[j])
				j++
			}
		}
		if i < m {
			res = append(res, left[i:]...)
		}
		if j < n {
			res = append(res, right[j:]...)
		}
		return res
	}

	mid := n / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return _merge(left, right)
}

// 合并2个有序数组
func merge(left, right []int) []int {
	llen, rlen := len(left), len(right)
	res := make([]int, 0, llen+rlen)

	var i, j int
	for i < llen && j < rlen {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	// 左边数组还有元素没遍历完
	if i < llen {
		res = append(res, left[i:]...)
	}
	if j < rlen {
		res = append(res, right[j:]...)
	}
	return res
}
