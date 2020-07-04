package main

func oneEditAway(first string, second string) bool {
	size1, size2 := len(first), len(second)
	if size1 > size2+1 || size2 > size1+1 {
		return false
	}

	left := 0
	for {
		if left >= size1 || left >= size2 || first[left] != second[left] {
			break
		}
		left++
	}

	// fmt.Println(left, size1, size2)
	if left == size1 && left == size2 {
		return true
	}
	right1, right2 := size1-1, size2-1
	for {
		if right1 < 0 || right2 < 0 || first[right1] != second[right2] {
			break
		}
		right1--
		right2--
	}

	return (right1 >= right2 && right1 == left) || (right1 < right2 && right2 == left)
}

// func main() {
// 	fmt.Println(oneEditAway("pale", "ple"))
// 	fmt.Println(oneEditAway("pales", "pal"))
// 	fmt.Println(oneEditAway("", "p"))
// 	fmt.Println(oneEditAway("teacher", "bleacher"))
// 	fmt.Println(oneEditAway("teacher", "leacher"))
// }
