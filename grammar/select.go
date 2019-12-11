package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	str_chan := make(chan string, 1)

	int_chan <- 1
	str_chan <- "Hello"

	select {
	case val := <-int_chan:
		fmt.Println(val)
	case val := <-str_chan:
		panic(val)
	}
}

// 随机执行，可能panic也可能输出1，单个chan如果无缓冲时将会阻塞，select执行原则
// select 中只要有一个case能return，则立刻执行。
// 当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。
// 如果没有一个case能return则可以执行”default”块
