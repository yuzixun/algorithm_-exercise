package main

func findTheLongestSubstring(s string) int {
	cache := map[rune]int{'a': 1, 'e': 2, 'i': 4, 'o': 8, 'u': 16}

	helper := map[int]int{0: 0}

	result := 0
	flag := 0
	for i, item := range s {
		flag ^= cache[item]

		v, ok := helper[flag]
		if !ok {
			helper[flag] = i + 1
		}
		// fmt.Println(helper, i, v)
		result = max(result, i+1-helper[flag])
	}
	// fmt.Println(result)
	return result
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	findTheLongestSubstring("eleetminicoworoep")
	findTheLongestSubstring("leetcodeisgreat")
	findTheLongestSubstring("bcbcbc")
}
