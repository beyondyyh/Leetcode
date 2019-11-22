package main

import "fmt"

func fibnacci(N int) int {
	x, y := 0, 1
	for i := 0; i < N; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	fmt.Println(fibnacci(10)) // 55
}
