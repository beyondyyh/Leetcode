package main

import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s, "len:", len(s), "cap:", cap(s))

	/*
		var a map[string]string
		fmt.Println(make(map[string]string))
		fmt.Println(new(map[string]string))
		fmt.Println(a)
	*/

	// 两个考点，1. make的默认值，2. slice动态扩容
	// 1. make()一个slice可以有三个参数，1是类型 2是length，3是cap，第三个参数没指定时默认是cap=length
	// 2. slice动态扩容，if len<1024时，成倍增加，此题默认cap=5，并且生成了5个默认值0，append时cap不够扩容成10，在后面追加1，2，3
	// 所以为[0 0 0 0 0 1 2 3] len: 8 cap: 10
	// make初始化是由默认值的哦，此处默认值为0

	// 正确的做法：len设置为0
	ss := make([]int, 0)
	ss = append(ss, 1, 2, 3)
	fmt.Println(ss, "len:", len(ss), "cap:", cap(ss)) // [1 2 3] len: 3 cap: 4
	fmt.Println(make([]int, 5), new([]int), new(map[string]string))

	/*
		// 考点：new是指针操作，无法append，panic
		s1 := new([]int)
		s1 = append(s1, 1)
		fmt.Println(s1)
	*/
}
