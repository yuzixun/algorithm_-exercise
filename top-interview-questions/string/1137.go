package main

import "fmt"

func partition(s string) [][]string {
	cur := []string{}
	result := [][]string{}
	helper(s, 0, len(s), cur, &result)
	fmt.Println(result)
	return result
}

func check(str string) bool {
	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
func helper(s string, left, size int, cur []string, result *[][]string) {
	if left == size {
		new := make([]string, len(cur))
		copy(new, cur)
		// fmt.Println(new, cur)
		*result = append(*result, new)
		return
	}

	for i := left; i < size; i++ {
		if check(s[left : i+1]) {
			helper(s, i+1, size, append(cur, s[left:i+1]), result)
		}
	}
}

func main() {
	partition("aab")
	partition("cbbbcc")
	partition("ababbbabbaba")
}
