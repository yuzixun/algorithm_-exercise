package main

func groupAnagrams(strs []string) [][]string {
	hash := make(map[[26]int][]string)

	for i := 0; i < len(strs); i++ {
		var letters [26]int

		for j := 0; j < len(strs[i]); j++ {
			letters[strs[i][j]-'a']++
		}

		hash[letters] = append(hash[letters], strs[i])
	}

	result, count := make([][]string, len(hash)), 0
	for _, v := range hash {
		result[count] = v
		count++
	}
	// fmt.Println(result)
	return result
}

// func main() {
// 	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
// }
