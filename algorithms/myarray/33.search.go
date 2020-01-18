package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	L := len(nums)
	left, right := 0, L-1
	for left <= right {
		mid := (left + right) / 2
		fmt.Printf("left:%d right:%d mid:%d nums[mid]:%d\n", left, right, mid, nums[mid])
		if target == nums[mid] {
			return mid
		} else if nums[mid] < nums[right] {
			// 右半段是有序的
			if target > nums[mid] && target <= nums[right] {
				fmt.Println("a")
				left = mid + 1
			} else {
				fmt.Println("b")
				right = mid - 1
			}
		} else {
			// 左半段是连续的
			if target >= nums[left] && target < nums[mid] {
				fmt.Println("c")
				right = mid - 1
			} else {
				fmt.Println("d")
				left = mid + 1
			}
		}
	}
	return -1
}
