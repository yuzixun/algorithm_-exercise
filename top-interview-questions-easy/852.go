package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1

	var mid int
	for left <= right {
		mid = left + (right-left)/2
		if letters[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// fmt.Println(left)
	return letters[left%len(letters)]
}

func main() {
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'))
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd'))
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'g'))
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'j'))
	fmt.Println(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'k'))
}
