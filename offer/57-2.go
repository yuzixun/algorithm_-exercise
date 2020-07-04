package main

func findContinuousSequence(target int) [][]int {
	result := [][]int{}

	if target <= 2 {
		return result
	}

	left, right := 1, 2
	max := target/2 + 1
	sum := left + right
	for left <= max && right <= max {
		// fmt.Println(left, right, sum)
		if sum > target {
			sum -= left
			left++
		} else if sum < target {
			right++
			sum += right
		} else {
			cur := []int{}
			for i := left; i <= right; i++ {
				cur = append(cur, i)
			}
			result = append(result, cur)
			sum -= left
			left++
		}
	}
	// fmt.Println(result)
	return result
}

// func main() {
// 	findContinuousSequence(9)
// 	findContinuousSequence(15)
// }
