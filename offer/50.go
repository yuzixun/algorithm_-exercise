package main

import "fmt"

func firstUniqChar(s string) byte {
	hash := make([]int, 26)

	for _, i := range s {
		fmt.Println(i - 'a')
		hash[i-'a']++
	}

	for _, i := range s {
		if hash[i-'a'] == 1 {
			return byte(i)
		}
	}
	// for i := 0; i < len(s); i++ {
	// 	hash[s[i]]++
	// }

	// for i := 0; i < len(s); i++ {
	// 	if hash[s[i]] == 1 {
	// 		return s[i]
	// 	}
	// }
	return ' '
}

// func main() {
// 	fmt.Println(firstUniqChar("abaccdeff") == 'b')
// 	fmt.Println(firstUniqChar(""))

// }
