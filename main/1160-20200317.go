package main

func countCharacters(words []string, chars string) int {
	result := 0
	arr := make([]int, 26)
	cur := make([]int, 26)
	for i := 0; i < len(chars); i++ {
		arr[chars[i]-'a']++
	}

	for _, word := range words {
		copy(cur, arr)
		curCount := 0
		for i := 0; i < len(word); i++ {
			cur[word[i]-'a']--
			curCount++
			if cur[word[i]-'a'] < 0 {
				curCount = 0
				break
			}
		}
		result += curCount
	}
	return result
}

func main() {
	countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach")
	countCharacters([]string{"hello", "world", "leetcode"}, "welldonehoneyr")
}
