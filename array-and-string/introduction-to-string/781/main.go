package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	size := len(strs)
	if size == 0 {
		return ""
	} else if size == 1 {
		return strs[0]
	}

	result := strs[0]
	for i := 1; i < size; i++ {
		j := 0
		for ; j < len(result) && j < len(strs[i]); j++ {
			if strs[i][j] != result[j] {
				break
			}
		}

		result = result[:j]
	}

	return result
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
