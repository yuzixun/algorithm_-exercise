package main

import "fmt"

func maxDepthAfterSplit(seq string) []int {
	result := make([]int, len(seq))
	flag := 1

	for i := 0; i < len(seq); i++ {
		if seq[i] == '(' {
			flag ^= 1
			result[i] = flag
		} else {
			result[i] = flag
			flag ^= 1
		}
	}

	return result
}

func main() {
	fmt.Println(maxDepthAfterSplit("(()())"))
	fmt.Println(maxDepthAfterSplit("()(())()"))
}
