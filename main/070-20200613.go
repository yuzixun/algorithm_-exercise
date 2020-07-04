package main

import "fmt"

func climbStairs(n int) int {
	pre, cur := 1, 2
	for i := 3; i <= n; i++ {
		cur, pre = cur+pre, cur
	}
	return cur
}

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(19))
}
