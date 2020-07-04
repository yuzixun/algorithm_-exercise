package main

func maxProfit(prices []int) int {
	size, result := len(prices), 0
	if size == 0 {
		return 0
	}
	minPrice := prices[0]

	for i := 1; i < size; i++ {
		result = max(result, prices[i]-minPrice)
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	// fmt.Println(result)
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxProfit([]int{7, 1, 5, 3, 6, 4})
	maxProfit([]int{7, 6, 4, 3, 1})
}
