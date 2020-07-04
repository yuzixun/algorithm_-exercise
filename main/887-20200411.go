package main

// func superEggDrop(K int, N int) int {
// 	dp := make([][]int, K+1)
// 	for i := 0; i < len(dp); i++ {
// 		dp[i] = make([]int, N+1)
// 	}

// 	m := 0
// 	for dp[K][m] < N {
// 		m++
// 		for k := 1; k <= K; k++ {
// 			dp[k][m] = dp[k][m-1] + dp[k-1][m-1] + 1

// 		}
// 		// fmt.Println(m, N, dp)
// 	}

// 	// fmt.Println(m)
// 	return m
// }

func superEggDrop(K int, N int) int {
	times := 1
	for {
		if getTimes(times, K) >= N {
			break
		}
		times++
	}
	return times
}
func getTimes(times, eggCount int) int {
	if times == 1 || eggCount == 1 {
		return times
	}

	return getTimes(times-1, eggCount) + 1 + getTimes(times-1, eggCount-1)
}

// func superEggDrop(K int, N int) int {
// 	memo := make(map[int]map[int]int, K)
// 	MAX := 1 << 32
// 	var helper func(k, n int) int
// 	helper = func(k, n int) int {
// 		if k == 1 {
// 			return n
// 		}

// 		if n == 0 {
// 			return 0
// 		}

// 		if v, ok := memo[k][n]; ok {
// 			return v
// 		}

// 		res := MAX
// 		left, right := 1, n
// 		for left <= right {
// 			mid := left + (right-left)/2
// 			broken := helper(k-1, mid-1)
// 			notBroken := helper(k, n-mid)
// 			if broken > notBroken {
// 				right = mid - 1
// 				res = min(res, broken+1)
// 			} else {
// 				left = mid + 1
// 				res = min(res, notBroken+1)
// 			}
// 		}

// 		if _, ok := memo[k]; !ok {
// 			memo[k] = make(map[int]int)
// 		}
// 		memo[k][n] = res
// 		// if _, ok := memo[k]; ok {

// 		// } else {
// 		// 	memo[k] = map[int]int{n: res}
// 		// }

// 		return res
// 	}

// 	return helper(K, N)
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func main() {
	superEggDrop(1, 2)
	superEggDrop(2, 6)
	superEggDrop(3, 14)
	superEggDrop(4, 5000)
}
