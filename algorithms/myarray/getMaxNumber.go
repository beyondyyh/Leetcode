package myarray

import "fmt"

// func main() {
// 	arr := []int{1, 2, 2, 3, 4, 5, 10, 9, 8, 3, 3, 3, 3, 2, 1}
// 	fmt.Println(getMax(arr)) // 10
// }

func getMax(arr []int) int {
	l := len(arr)
	if l == 1 {
		return arr[0]
	}

	mid := l / 2
	fmt.Printf("len:%d mid:%d arr:%v \n", l, mid, arr)
	if arr[mid-1] > arr[mid] {
		return getMax(arr[:mid])
	} else if arr[mid-1] < arr[mid] {
		return getMax(arr[mid:])
	} else {
		left := getMax(arr[:mid])
		right := getMax(arr[mid:])
		if left > right {
			return left
		}
		return right
	}
}

// output:
// len:15 mid:7 arr:[1 2 2 3 4 5 10 9 8 3 3 3 3 2 1]
// len:7 mid:3 arr:[1 2 2 3 4 5 10]
// len:4 mid:2 arr:[3 4 5 10]
// len:2 mid:1 arr:[5 10]
