package main

import "fmt"

func findAnagrams(s string, p string) []int {
	// if len(s) < len(p) {
	// 	return 0
	// }
	var ret []int
	sArray := []rune(s)
	pArray := []rune(p)
	var pMap map[rune]int = make(map[rune]int)
	for _, r := range pArray {
		pMap[r] += 1
	}
	fmt.Printf("pmap is %v\n", pMap)
	pSize := len(p)

	begin, end := 0, 0
	for end < len(s) {
		v := sArray[end]
		fmt.Printf("v, %d\n", v)
		pVal, ok := pMap[v]
		if ok && pVal > 0 {
			pSize -= 1
		}
		pMap[v] -= 1

		if end-begin == len(p)-1 {
			if pSize == 0 {
				ret = append(ret, begin)
			}
			// fmt.Println(begin)
			temp := sArray[begin]
			if pVal, ok := pMap[temp]; ok && pVal >= 0 {
				fmt.Printf("pval, %v, %d, %T\n", temp, pMap[temp], pMap[temp])

				pSize += 1
			}
			pMap[temp] += 1

			begin++
		}
		end++
	}

	fmt.Println(ret)
	return ret
}

// func main() {
// 	// findAnagrams("cbaebabacd", "abc")
// 	findAnagrams("abab", "ab")
// }
