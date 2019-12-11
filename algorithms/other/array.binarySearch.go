package main

import "fmt"

// 二分查找
func binarySearch(nums []int, target int) bool {
	l := len(nums)
	left, right := 0, l-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return false
}

func main() {
	nums := []int{-1, 0, 1, 2, 3, 4, 5, 10, 11, 19}
	fmt.Println(binarySearch(nums, 1))
}
