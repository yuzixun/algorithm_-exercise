package main

func lexicalOrder(n int) []int {
	result := []int{}
	helper(&result, -1, n)
	// fmt.Println(result)
	return result
}

func helper(result *[]int, current int, maxNum int) {
	if current != -1 && current > maxNum {
		return
	}

	if current != -1 {
		*result = append(*result, current)
	}

	for i := 0; i <= 9; i++ {
		if current == -1 {
			if i == 0 {
				continue
			} else {
				current = 0
			}
		}

		helper(result, current*10+i, maxNum)
	}
}

func main() {
	lexicalOrder(13)
}
