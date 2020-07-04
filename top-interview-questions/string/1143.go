package main

import "fmt"

func firstUniqChar(s string) int {
	arr := [26]int{}

	for _, item := range s {
		arr[item-'a']++
	}

	for i, item := range s {
		if arr[item-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}
