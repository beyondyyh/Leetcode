package main

import "fmt"

func main() {
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))
	// 4 1 3
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	// return t
	return
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
