package main

import "fmt"

func canPermutePalindrome(s string) bool {
	size := len(s)
	hash := map[rune]int{}
	for _, item := range s {
		hash[item]++
	}

	if size%2 == 1 {
		countOne := 0
		for _, value := range hash {
			if value > 2 {
				return false
			} else if value == 1 {
				countOne++
				if countOne > 1 {
					return false
				}
			}
		}
	} else {
		for _, value := range hash {
			if value != 2 {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(canPermutePalindrome("tactcoa"))
}
