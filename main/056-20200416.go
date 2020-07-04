package main

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	size := len(intervals)
	if size == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return (intervals[i][0] < intervals[j][0]) || (intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1])
	})

	result := [][]int{}
	result = append(result, intervals[0])

	for i := 1; i < size; i++ {
		if intervals[i][0] <= result[len(result)-1][1] {
			if intervals[i][1] > result[len(result)-1][1] {
				result[len(result)-1][1] = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
		}
	}

	// fmt.Println(result)
	return result
}

func main() {
	merge([][]int{
		[]int{1, 3},
		[]int{2, 6},
		[]int{8, 10},
		[]int{15, 18},
	})

	merge([][]int{
		[]int{1, 4},
		[]int{4, 5},
	})

	merge([][]int{
		[]int{4, 5},
		[]int{1, 4},
		[]int{0, 1},
	})

	merge([][]int{
		[]int{1, 4},
		[]int{0, 0},
	})
	merge([][]int{
		[]int{1, 4},
		[]int{0, 2},
	})
	merge([][]int{
		[]int{1, 4},
		[]int{1, 5},
	})

	merge([][]int{
		[]int{1, 4},
		[]int{2, 3},
	})
}
