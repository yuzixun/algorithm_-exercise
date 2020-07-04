package main

func sumFourDivisors(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		result, count := 0, 0
		for j := 1; j*j <= nums[i]; j++ {
			if nums[i]%j == 0 {
				result += j
				count++
				if j*j < nums[i] {
					result += nums[i] / j
					count++
				}
			}
			if count > 4 {
				break
			}
		}

		if count == 4 {
			sum += result
		}
	}
	// fmt.Println(sum)
	return sum
}

// func main() {
// 	sumFourDivisors([]int{21, 4, 7})
// }
