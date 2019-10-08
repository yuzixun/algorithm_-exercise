package main

import (
	"math"
)

func maxProfit(prices []int) int {
	var profit int
	size := len(prices)
	for i := 0; i < size-1; i++ {
		delta := (prices[i+1] - prices[i])
		profit += int(math.Max(0, float64(delta)))
	}
	return profit
}
