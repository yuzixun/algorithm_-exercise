package main

import (
	"sort"
	"strconv"
	"strings"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return s[i]+s[j] < s[j]+s[i]
}

func minNumber(nums []int) string {
	strs := []string{}
	for i := 0; i < len(nums); i++ {
		strs = append(strs, strconv.Itoa(nums[i]))
	}
	sort.Sort(byLength(strs))
	return strings.Join(strs, "")
}

// func main() {
// 	// fmt.Println("peach" > "babab")
// 	// fruits := []string{"peach", "banana", "kiwi"}
// 	// sort.Sort(byLength(fruits))
// 	// fmt.Println(fruits)
// 	fmt.Println(minNumber([]int{10, 2}))
// 	fmt.Println(minNumber([]int{3, 30, 34, 5, 9}))
// }
