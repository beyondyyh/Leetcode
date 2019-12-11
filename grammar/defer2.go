package main

import (
	"fmt"
)

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	defer calc("2", a, calc("20", a, b))
}

// defer先入后出原则
// index1 肯定是最后执行，但是index1的第三个参数是个函数，所以按照顺序最先执行；到index2时跟index1一样，限制性第三个参数即index20
// 所以执行顺序： index10->index20->index2->index1
// output:
// 10 1 2 3
// 20 1 2 3
// 2 1 3 4
// 1 1 3 4
