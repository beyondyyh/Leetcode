package myarray

// 双指针法
func maxArea(height []int) int {
	maxarea, left, right := 0, 0, len(height)-1
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		maxarea = max(maxarea, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxarea
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
