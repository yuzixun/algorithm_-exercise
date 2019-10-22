package main

func maxProfit(k int, prices []int) int {
	if len(prices) == 0 || k == 0 {
		return 0
	}
	minInt := -1 << 32
	profits := make([][][]int, len(prices))
	for ii := 0; ii < len(prices); ii++ {
		arr := make([][]int, k+1)
		for kk := 0; kk <= k; kk++ {
			arr[kk] = []int{minInt, minInt}
		}
		profits[ii] = arr
	}

	profits[0][0] = []int{0, -prices[0]}
	size := len(prices)
	for ii := 1; ii < size; ii++ {
		for kk := 0; kk <= k; kk++ {
			if kk == 0 {
				profits[ii][0][0] = profits[ii-1][0][0]
				profits[ii][0][1] = max(profits[ii-1][0][1], profits[ii-1][0][0]-prices[ii])
			} else {
				profits[ii][kk][0] = max(profits[ii-1][kk][0], profits[ii-1][kk-1][1]+prices[ii])
				profits[ii][kk][1] = max(profits[ii-1][kk][1], profits[ii-1][kk][0]-prices[ii])
			}
		}
	}

	max := -1
	for _, arr := range profits[size-1] {
		if max < arr[0] {
			max = arr[0]
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	// prices := []int{2, 4, 1}
// 	prices := []int{3, 2, 6, 5, 0, 3}
// 	k := 2
// 	fmt.Println(maxProfit(k, prices))
// }
