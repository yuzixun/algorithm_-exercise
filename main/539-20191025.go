package main

import (
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	size := len(timePoints)
	min := 24 * 60
	times := make([]int, size)

	for i, timePoint := range timePoints {
		time := strings.Split(timePoint, ":")
		h, _ := strconv.Atoi(time[0])
		m, _ := strconv.Atoi(time[1])
		times[i] = h*60 + m
	}
	sort.Ints(times)

	for i := 1; i < size; i++ {
		interval := times[i] - times[i-1]
		if min > interval {
			min = interval
		}
	}
	back := (times[0] + 1440 - times[size-1])
	if min > back {
		min = back
	}
	return min
}

// func main() {
// 	timePoints := []string{"23:59", "00:00"}

// 	fmt.Println(findMinDifference(timePoints))
// }
