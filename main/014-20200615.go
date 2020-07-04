package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	result := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0
		for ; j < len(strs[i]) && j < len(result); j++ {

			if result[j] != strs[i][j] {
				break
			}
		}

		result = result[0:j]
		if result == "" {
			return ""
		}
	}

	return result
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "dacecar", "dar"}))
	fmt.Println(longestCommonPrefix([]string{"d", "a"}))
	fmt.Println(longestCommonPrefix([]string{"aa", "a"}))
	fmt.Println(longestCommonPrefix([]string{"aa"}))
	fmt.Println(longestCommonPrefix([]string{"aca", "cba"}))
}
