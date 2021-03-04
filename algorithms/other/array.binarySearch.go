package other

// 二分查找
func binarySearch(nums []int, target int) bool {
	l := len(nums)
	left, right := 0, l-1
	for left <= right {
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
