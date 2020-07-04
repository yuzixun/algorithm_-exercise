package main

import "fmt"

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// result := false
	// result = result || rec1[0] < rec2[0] && rec2[0] < rec1[2] && rec1[1] < rec2[1] && rec2[1] < rec1[3]
	// result = result || rec1[0] < rec2[2] && rec2[2] < rec1[2] && rec1[1] < rec2[3] && rec2[3] < rec1[3]
	// result = result || rec2[0] < rec1[0] && rec1[0] < rec2[2] && rec2[1] < rec1[1] && rec1[1] < rec2[3]
	// return result

	return (rec2[0]-rec1[2])*(rec2[2]-rec1[0]) < 0 && (rec2[1]-rec1[3])*(rec2[3]-rec1[1]) < 0
}

func main() {
	fmt.Println(isRectangleOverlap([]int{0, 0, 2, 2}, []int{1, 1, 3, 3}))
	fmt.Println(isRectangleOverlap([]int{0, 0, 1, 1}, []int{1, 0, 2, 1}))
}
