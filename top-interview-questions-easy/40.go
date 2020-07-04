package main

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
		for ; j < len(strs[i]) && j < len(result); j++ {
			if strs[i][j] != result[j] {
				break
			}
		}
		result = result[:j]
	}
	return result
}

func main() {
	longestCommonPrefix([]string{"flower", "flow", "flight"})
	longestCommonPrefix([]string{"dog", "racecar", "car"})
	longestCommonPrefix([]string{"aa", "a"})
	longestCommonPrefix([]string{"a", "b"})
	longestCommonPrefix([]string{"aca", "cba"})
	longestCommonPrefix([]string{"dog", "racecar", "car"})
}
