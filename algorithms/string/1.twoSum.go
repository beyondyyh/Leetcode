package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for k, v := range nums {
		diff := target - v
		if kk, exists := hash[diff]; exists {
			return []int{kk, k}
		}
		hash[v] = k
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
