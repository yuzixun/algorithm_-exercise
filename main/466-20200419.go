package main

// func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
// 	count2, index := 0, 0
// 	for i := 0; i < n1; i++ {
// 		for j := 0; j < len(s1); j++ {

// 			if s1[j] == s2[index] {
// 				if index == len(s2)-1 {
// 					index = 0
// 					count2++
// 				} else {
// 					index++
// 				}
// 			}
// 		}
// 	}

// 	// fmt.Println(count2, n2)
// 	return count2 / n2
// }

func main() {
	// getMaxRepetitions("acb", 4, "ab", 2)
	getMaxRepetitions("abaacdbac", 100, "adcbd", 4)
}
