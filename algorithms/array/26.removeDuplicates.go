package main

import "fmt"

func removeDuplicates(nums []int) int {
	L := len(nums)
	if L == 0 {
		return 0
	}

	fmt.Printf("before:%v\n", nums)
	i := 0
	fmt.Printf("%v ", nums[i])
	for j := 1; j < L; j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
			fmt.Printf("%v ", nums[i])
		}
	}
	fmt.Println("\n********")
	return i + 1
}
