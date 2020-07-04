package main

func nthUglyNumber(n int) int {
	result := []int{1}
	// var arr2, arr3, arr5 []int
	var i, j, k, cur int

	for index := 1; index < n; index++ {
		t2 := result[i] * 2
		t3 := result[i] * 3
		t5 := result[i] * 5
		cur = min(min(t2, t3), t5)
		result = append(result, cur)
		// fmt.Println(result, i, j, k)
		if t2 == cur {
			i++
		}
		if t3 == cur {
			j++
		}
		if t5 == cur {
			k++
		}
	}

	// fmt.Println(result, result[n-1])
	return result[n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	nthUglyNumber(10)
}
