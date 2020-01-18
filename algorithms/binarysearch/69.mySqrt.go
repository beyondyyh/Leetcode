package binarysearch

import _ "fmt"

func mySqrt(x int) int {
	left, right := 0, x
	for left < right {
		// >> 右移，高位补符号位 右移位表示除2
		// << 左移，左移位表示乘2
		// 每次mid取值为(l + r + 1) / 2
		// 这是为了防止当l = r - 1时，出现死循环的情况
		mid := (left + right + 1) >> 1
		// fmt.Printf("l:%d r:%d mid:%d\n", left, right, mid)
		if mid*mid <= x {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}
