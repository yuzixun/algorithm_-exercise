package main

func findContinuousSequence(target int) [][]int {
	result := [][]int{}
	if target <= 2 {
		return result
	}
	left, right, max := 1, 2, target/2+1
	sum := left + right
	for left <= max && right <= max {
		if sum < target {
			right++
			sum += right
		} else if sum > target {
			sum -= left
			left++
		} else {
			temp := []int{}
			for i := left; i <= right; i++ {
				temp = append(temp, i)
			}
			result = append(result, temp)
			sum -= left
			left++
		}
	}

	// fmt.Println(result)
	return result
}

func main() {
	findContinuousSequence(9)
	findContinuousSequence(15)
}
