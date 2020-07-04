package main

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 {
		return true
	}

	var temp []int
	ppIndex := 0
	for i := 0; i < len(pushed); i++ {
		// fmt.Println("temp is ", temp)
		if pushed[i] != popped[ppIndex] {
			temp = append(temp, pushed[i])
		} else {
			ppIndex++

			j := len(temp) - 1
			for ; j >= 0; j-- {
				// fmt.Println(temp[j], popped[ppIndex], j, ppIndex)
				if temp[j] != popped[ppIndex] {
					break
				}
				ppIndex++
				temp = temp[:j]
			}
			// fmt.Println("j is ", j, temp)
		}
	}

	return len(temp) == 0
}

// func main() {
// 	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
// 	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}))
// 	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}))
// }
