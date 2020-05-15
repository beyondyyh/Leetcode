package main

import (
	"fmt"
	"time"
)

// 生产者：生成 factor 整数倍的序列
func producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// 消费者
func consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int, 64)

	go producer(3, ch)
	// go producer(5, ch)
	go producer(7, ch)
	go consumer(ch)

	time.Sleep(500 * time.Millisecond)

	// // Ctrl+C 退出
	// sig := make(chan os.Signal, 1)
	// signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	// fmt.Printf("quit (%v)\n", <-sig)
}
