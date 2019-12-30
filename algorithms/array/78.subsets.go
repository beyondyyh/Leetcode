package main

import "fmt"

func subsets(nums []int) [][]int {
	res := make([][]int, 1)
	fmt.Println(res)
	for _, num := range nums {
		for _, sub := range res {
			clone := make([]int, len(sub), len(sub)+1)
			copy(clone, sub)
			clone = append(clone, num)
			res = append(res, clone)
		}
	}
	return res
}
