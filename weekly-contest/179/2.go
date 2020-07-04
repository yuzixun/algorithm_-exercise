package main

func numTimesAllBlue(light []int) int {

	result, max := 0, 0
	for i := 0; i < len(light); i++ {

		if light[i] > max {
			max = light[i]
		}
		if max == i+1 {
			result++
		}
	}
	// fmt.Println(result)
	return result
}

func main() {
	numTimesAllBlue([]int{})
	numTimesAllBlue([]int{2, 1, 3, 5, 4})
	numTimesAllBlue([]int{3, 2, 4, 1, 5})
	numTimesAllBlue([]int{4, 1, 2, 3})
	numTimesAllBlue([]int{2, 1, 4, 3, 6, 5})
	numTimesAllBlue([]int{1, 2, 3, 4, 5, 6})
}
