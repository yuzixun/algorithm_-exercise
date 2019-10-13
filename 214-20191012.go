package main

func shortestPalindrome(s string) string {
	doubleS := s + "#" + reverseString(s)
	table := getTable(doubleS)
	return reverseString(s[table[len(table)-1]:]) + s
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func getTable(s string) []int {
	size := len(s)
	result := make([]int, size)

	j := 0
	i := 1
	for i < size {
		if s[i] == s[j] {
			j++
			result[i] = j
		} else {
			if j > 0 {
				j = result[j-1]
				continue
			}
		}

		i++
	}

	return result
}
