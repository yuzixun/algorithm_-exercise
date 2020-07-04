package main

func letterCombinations(digits string) []string {
	var result []string
	if digits == "" {
		return result
	}
	hash := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}

	var helper func(tmp string, index int)
	helper = func(tmp string, index int) {
		if index == len(digits) {
			result = append(result, tmp)
			return
		}

		letters := hash[digits[index]]

		for i := 0; i < len(letters); i++ {
			helper(tmp+string([]byte{letters[i]}), index+1)
		}
	}

	helper("", 0)
	// fmt.Println(result)
	return result
}

// func main() {
// 	letterCombinations("23")
// }
