package main

// func lengthOfLongestSubstring(s string) int {
// 	size := len(s)
// 	if size == 0 {
// 		return 0
// 	}

// 	slow, fast, result := 0, 0, 0
// 	arr := []byte{}
// 	for ; fast < size; fast++ {
// 		cur := s[fast]
// 		pos := 0
// 		for i, item := range arr {
// 			if item == cur {
// 				slow += i
// 				pos = i + 1
// 			}
// 		}

// 		arr = arr[pos:]
// 		arr = append(arr, cur)
// 		// fmt.Println("arr is ", arr, pos)
// 		if len(arr) > result {
// 			result = len(arr)
// 		}
// 	}

// 	// fmt.Println(result)
// 	return result
// }

func lengthOfLongestSubstring(s string) int {
	size := len(s)

	arr := make([]int, 256)
	left, right, result := 0, 0, 0
	for left < size {

		if right < size && arr[s[right]] == 0 {
			arr[s[right]]++
			right++
		} else {
			arr[s[left]]--
			left++
		}

		result = max(result, right-left)
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
	lengthOfLongestSubstring("abcabcbb")
	lengthOfLongestSubstring("bbbbb")
	lengthOfLongestSubstring("")
	lengthOfLongestSubstring("pwwkew")
	lengthOfLongestSubstring("aab")
}
