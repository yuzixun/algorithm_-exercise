package main

func permutation(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	if len(s) == 1 {
		return []string{s}
	}

	result := []string{}
	usedCache := []byte{}
	for i := 0; i < len(s); i++ {
		goToNext := false
		for k := 0; k < len(usedCache); k++ {
			if usedCache[k] == s[i] {
				goToNext = true
			}
		}
		if goToNext {
			continue
		} else {
			usedCache = append(usedCache, s[i])
		}
		// fmt.Println(s[:i] + s[i+1:])
		temp := permutation(s[:i] + s[i+1:])
		for j := 0; j < len(temp); j++ {
			result = append(result, string(s[i])+temp[j])
		}
	}

	// fmt.Println(result)
	return result
}

// func main() {
// 	permutation("abc")
// 	permutation("aab")
// }
