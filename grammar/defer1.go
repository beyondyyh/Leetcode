package main

import (
	"fmt"
)

func main() {
	defer_call()
}

// defer执行顺序是后入先出
// 是一个栈
// 最先打印正常输出，然后打印defer栈（先进后出），最后执行panic
func defer_call() {
	fmt.Println("aaaaa")
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

// output
// 打印后
// 打印中
// 打印前
// panic: 触发异常
