package main

import "fmt"

func majorityElementVote(nums []int) int {
	cur, count := 0, 0
	for _, v := range nums {
		fmt.Println(cur, count, v)
		if count == 0 {
			cur = v
		}

		if cur == v {
			count++
		} else {
			count--
		}
		fmt.Println("after :", cur, count, v)
	}
	return cur
}

func main() {
	majorityElementVote([]int{2, 2, 1, 1, 1, 2, 2})
	// majorityElementVote([]int{3, 2, 3})
}
