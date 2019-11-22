package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return nil
	}

	arr := make([][]int, 0)
	arr = append(arr, []int{1})
	for i := 1; i < numRows; i++ {
		tmp := []int{1}
		for j := 1; j < i; j++ {
			tmp = append(tmp, arr[i-1][j-1]+arr[i-1][j])
		}
		tmp = append(tmp, 1)
		arr = append(arr, tmp)
	}
	return arr
}

func main() {
	fmt.Println(generate(5))
	// out: [[1] [1 1] [1 2 1] [1 3 3 1] [1 4 6 4 1]]
}
