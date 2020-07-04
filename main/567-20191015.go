package main

func checkInclusion(s1 string, s2 string) bool {
	counter := make(map[rune]int)
	for _, l := range s1 {
		counter[l]++
	}

	size := len(s2)
	r2 := []rune(s2)
	for slow, fast := 0, 0; fast < size; fast++ {
		counter[r2[fast]]--
		if counter[r2[fast]] < 0 {
			for {
				l := r2[slow]
				slow++
				counter[l]++
				if counter[l] == 0 {
					break
				}
			}
		} else if fast-slow+1 == len(s1) {
			return true
		}
	}
	return len(counter) == 0
}

// func main() {
// 	s1, s2 := "adc", "dcda"

// 	// fmt.Println(checkInclusion(s1, s2))
// }
