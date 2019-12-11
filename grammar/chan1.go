package main

func main() {
	var ch1 = make(chan int)
	var ch2 = make(<-chan int)

	ch1 = ch2 // 报错：cannot use ch2 (type <-chan int) as type chan int in assignment
	ch2 = ch1 // ok
	_ = ch1
}
