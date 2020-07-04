package main

import (
	"math"
)

func twoSum(n int) []float64 {
	arr := [70]int{}

	for i := 1; i <= 6; i++ {
		arr[i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 6 * i; j >= i; j-- {
			arr[j] = 0
			for cur := 1; cur <= 6; cur++ {
				if j-cur < i-1 {
					break
				}
				arr[j] += arr[j-cur]
			}
		}
	}

	max := math.Pow(6, float64(n))
	var result []float64
	for i := n; i <= 6*n; i++ {
		cur := arr[i]
		result = append(result, float64(cur)/max)
	}

	// fmt.Println(result, arr)
	return result
}
