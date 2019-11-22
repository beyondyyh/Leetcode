package main

import (
	"fmt"
)

func main() {
	a1 := []int{1, 2, 3, 4, 5, 9}
	a2 := []int{1, 2, 5, 7, 9}
	a3 := []int{1, 3, 4, 5, 8, 9, 10}
	fmt.Println(arraysIntersection(a1, a2, a3)) // [1 5 9]
}

func arraysIntersection(nums1 []int, nums2 []int, nums3 []int) []int {
	var (
		x      = 0
		y      = 0
		z      = 0
		result []int
	)

	for x < len(nums1) && y < len(nums2) && z < len(nums3) {
		fmt.Printf("x:%d->%d y:%d->%d z:%d->%d\n", x, nums1[x], y, nums2[y], z, nums3[z])
		if nums1[x] == nums2[y] && nums2[y] == nums3[z] {
			result = append(result, nums1[x])
			x++
			y++
			z++
		} else if nums1[x] < nums2[y] || nums1[x] < nums3[z] {
			x++
		} else if nums2[y] < nums1[x] || nums2[y] < nums3[z] {
			y++
		} else if nums3[z] < nums1[x] || nums3[z] < nums2[y] {
			z++
		}
	}
	return result
}
