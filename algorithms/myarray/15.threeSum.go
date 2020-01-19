package myarray

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int
	length := len(nums)
	if length < 3 {
		return res
	}
	// 排序
	sort.Ints(nums)

	for i := 0; i < length; i++ {
		// 如果当前数字大于0，则三数之和一定大于0，所以结束循环
		if nums[i] > 0 {
			break
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		L, R := i+1, length-1
		// fmt.Printf("i:%d->%d L:%d->%d R:%d->%d sum:%d\n",
		//		i, nums[i], L, nums[L], R, nums[R], nums[i]+nums[L]+nums[R])
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			switch {
			case sum < 0:
				L++
			case sum > 0:
				R--
			default:
				res = append(res, []int{nums[i], nums[L], nums[R]})
				// 为避免重复添加，L 和 R 都需要移动到不同的元素上
				L, R = next(nums, L, R)
			}
		} // end inner for
	} // end outer for
	return res
}

func next(nums []int, L, R int) (int, int) {
	for L < R {
		switch {
		case nums[L] == nums[L+1]:
			L++
		case nums[R] == nums[R-1]:
			R--
		default:
			L++
			R--
			return L, R
		}
	}
	return L, R
}
