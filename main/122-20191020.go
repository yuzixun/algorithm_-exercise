package main

func maxProfit3(prices []int) int {
	var result int
	for i := 0; i < len(prices)-1; i++ {
		delta := (prices[i+1] - prices[i])
		if delta > 0 {
			result += delta
		}
	}
	return result
}

// func main() {
// 	// prices := []int{1, 2, 3, 4, 5}
// 	prices := []int{}

// 	fmt.Println(maxProfit(prices))
// }
