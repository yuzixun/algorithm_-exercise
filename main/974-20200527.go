package main

import "fmt"

func subarraysDivByK(A []int, K int) int {
	cache := map[int]int{}
	cache[0] = 1
	preSum, count := 0, 0

	for _, item := range A {
		preSum = (preSum + item) % K
		if preSum < 0 {
			preSum += K
		}
		count += cache[preSum]
		cache[preSum]++
	}

	return count
}

func main() {
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
}
