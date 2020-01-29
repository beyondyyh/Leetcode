package other

import "fmt"

// QuickSort 快排，会延伸很多变种
func QuickSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	mid := data[0]
	head, tail := 0, len(data)-1
	for i := 1; i <= tail; {
		fmt.Printf("i:%d, head:%d, tail:%d, mid:%d, data[i]:%d, data:%v\n",
			i, head, tail, mid, data[i], data)
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	// data[head] = mid
	QuickSort(data[:head])
	QuickSort(data[head+1:])
	return data
}

// func main() {
// 	nums := []int{8, 9, 1, 3, 2, 0, 7, 11}
// 	fmt.Printf("befort: %v\n", nums)
// 	fmt.Printf("after: %v\n", QuickSort(nums))
// }
