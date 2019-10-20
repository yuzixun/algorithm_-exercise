package main

import (
	"fmt"
)

func maxProfit1(prices []int) int {
	var maxProfit int
	min := 1 << 32
	fmt.Println(min)
	for _, price := range prices {
		if price < min {
			// fmt.Println("min is ", min)
			min = price
		}
		currentProfit := price - min
		// fmt.Println("price ", price, "currentProfit ", currentProfit, min, maxProfit)
		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}
	}
	return maxProfit
}
