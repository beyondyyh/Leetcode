package myarray

import "fmt"

func searchRange(nums []int, target int) []int {
	return []int{
		getLeftBound(nums, target),
		getRightBound(nums, target),
	}
}

func getLeftBound(nums []int, target int) int { // {{{
	L := len(nums)
	if L == 0 {
		return -1
	}

	// note, 搜索区间[left, right)
	left, right := 0, L

	for left < right {
		mid := (left + right) / 2
		fmt.Printf("left:%d right:%d mid:%d nums[mid]:%d\n", left, right, mid, nums[mid])
		// 展开是为了便于理解，可以合并
		if target == nums[mid] {
			// 当target==nums[mid]时，不要立即返回，而是减小「搜索区间」的下界 right，使得区间不断向左收缩，达到锁定左侧边界的目的
			right = mid // 注意，不是mid-1，左闭右开区间
		} else if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid // 注意，不是mid-1，左闭右开区间
		}
	}

	// 处理边界值
	// target比所有数都大
	if left == L {
		return -1
	}

	if target == nums[left] {
		return left
	}
	// target可能不存在于数组中
	return -1
} // }}}

func getRightBound(nums []int, target int) int { // {{{
	fmt.Println("***********")
	L := len(nums)
	if L == 0 {
		return -1
	}

	left, right := 0, L
	for left < right {
		mid := (left + right) / 2
		fmt.Printf("left:%d right:%d mid:%d nums[mid]:%d\n", left, right, mid, nums[mid])
		if target == nums[mid] {
			// 当target==nums[mid]时，不要立即返回，而是增大「搜索区间」的下界 left，使得区间不断向右收缩，达到锁定右侧边界的目的
			left = mid + 1
		} else if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid
		}
	}

	// 处理边界值
	// target比所有数都小
	if right == 0 {
		return -1
	}

	if target == nums[right-1] {
		return right - 1
	}
	// target可能不存在于数组中
	return -1
} // }}}
