package main

import (
	"strconv"
)

func translateNum(num int) int {
	var count int
	dfs(strconv.Itoa(num), &count)
	return count
}

func dfs(num string, count *int) {
	if num == "" {
		(*count)++
		return
	}
	dfs(num[1:], count)

	if num[0] == '1' || (num[0] == '2' && num[1] <= '5') {
		dfs(num[2:], count)
	}
}

func main() {
	translateNum(12258)
}
