package main

func longestPrefix(s string) string {
	result := ""
	size := len(s)
	for i := 0; i < size; i++ {
		if s[:i] == s[size-i:] {
			result = s[:i]
		}
	}
	// fmt.Println(result)
	return result
}

// func main() {
// 	longestPrefix("level")
// 	longestPrefix("ababab")
// 	longestPrefix("leetcodeleet")
// 	longestPrefix("a")
// 	longestPrefix("aaaaa")
// }
