package main

func minCostClimbingStairs(cost []int) int {
	size := len(cost)
	cache := make(map[int]int, size)
	cache[0] = 0
	cache[1] = 0
	return calc(cost, size, cache)
}

func calc(cost []int, step int, cache map[int]int) int {

	if val, ok := cache[step]; ok {
		return val
	}

	cache[step] = min(calc(cost, step-1, cache)+cost[step-1], calc(cost, step-2, cache)+cost[step-2])

	return cache[step]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// func main() {
// 	cost := []int{10, 15, 20}
// 	// cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
// 	fmt.Println(minCostClimbingStairs(cost))

// }
