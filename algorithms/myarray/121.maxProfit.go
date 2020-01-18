package myarray

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var min, max float64 = float64(prices[0]), 0
	for i := 1; i < len(prices); i++ {
		min = math.Min(min, float64(prices[i]))
		max = math.Max(max, float64(prices[i])-min)
	}

	fmt.Printf("min:%d max:%d\n", int(min), int(max))
	return int(max)
}
