package main

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	count := 0
	result := make([][3]int, n+1)
	var row, column int
	for _, seat := range reservedSeats {
		row, column = seat[0], seat[1]

		if column == 2 || column == 3 {
			result[row][0] = 1
		}

		if column == 4 || column == 5 {
			result[row][0] = 1
			result[row][1] = 1
		}
		if column == 6 || column == 7 {
			result[row][1] = 1
			result[row][2] = 1
		}
		if column == 8 || column == 9 {
			result[row][2] = 1
		}
	}

	var cur [3]int
	for i := 1; i < len(result); i++ {
		cur = result[i]
		if cur == [3]int{0, 0, 0} {
			count += 2
		} else if cur == [3]int{1, 1, 1} {
			continue
		} else {
			count++
		}
	}

	return count
}

func main() {
	maxNumberOfFamilies(3, [][]int{
		[]int{1, 2},
		[]int{1, 3},
		[]int{1, 8},
		[]int{2, 6},
		[]int{3, 1},
		[]int{3, 10},
	})
	maxNumberOfFamilies(2, [][]int{
		[]int{2, 1},
		[]int{1, 8},
		[]int{2, 6},
	})
	maxNumberOfFamilies(4, [][]int{[]int{4, 3},
		[]int{1, 4},
		[]int{4, 6},
		[]int{1, 7},
	})
}
