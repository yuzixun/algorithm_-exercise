package main

func wordBreak(s string, wordDict []string) bool {
	size := len(s)
	result := make([]bool, size+1)

	result[0] = true
	for i := 0; i <= size; i++ {
		for j := 0; j < i; j++ {
			if !result[j] {
				continue
			}

			for _, item := range wordDict {
				if s[j:i] == item {
					result[i] = true
					break
				}
			}
		}
	}
	// fmt.Println(result)
	return result[len(s)]
}

// func main() {
// 	wordBreak("leetcode", []string{"leet", "code"})
// 	wordBreak("applepenapple", []string{"apple", "pen"})
// 	wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
// }
