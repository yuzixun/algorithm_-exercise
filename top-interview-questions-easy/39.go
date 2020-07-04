package main

import "fmt"

func countAndSay(n int) string {
	result := "1"
	if n == 1 {
		return result
	}

	for i := 2; i <= n; i++ {
		count, cur := 1, result[0]
		temp := ""
		for j := 1; j < len(result); j++ {
			if cur != result[j] {
				// result = append(result, element)
				temp = fmt.Sprintf("%s%d%s", temp, count, string(cur))
				cur = result[j]
				count = 1
			} else {
				count++
			}
		}
		result = fmt.Sprintf("%s%d%s", temp, count, string(cur))
	}

	return result
}

func main() {
	for i := 0; i < 10; i++ {
		countAndSay(i)
	}
}
