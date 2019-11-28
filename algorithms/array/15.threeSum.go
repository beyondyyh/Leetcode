package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var res [][]int
	l := len(nums)
	if l < 3 {
		return res
	}
	// 排序
	sort.Ints(nums)

	for i := 0; i < l; i++ {
		// 如果当前数字大于0，则三数之和一定大于0，所以结束循环
		if nums[i] > 0 {
			break
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		L := i + 1
		R := l - 1
		fmt.Printf("i:%d->%d L:%d->%d R:%d->%d sum:%d\n",
			i, nums[i], L, nums[L], R, nums[R], nums[i]+nums[L]+nums[R])
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				var item []int
				item = append(item, nums[i], nums[L], nums[R])
				res = append(res, item)
				// 去重
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if sum < 0 {
				L++
			} else if sum > 0 {
				R--
			}
		} // endif
	} // endfor
	return res
}
