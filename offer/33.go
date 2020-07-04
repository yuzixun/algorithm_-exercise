package main

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 2 {
		return true
	}

	return verifyHelper(postorder, 0, len(postorder)-1)
}

func verifyHelper(postorder []int, left, right int) bool {
	if left >= right {
		return true
	}

	var i int
	for i = left; i < right; i++ {
		if postorder[i] > postorder[right] {
			break
		}
	}

	for j := i; j < right; j++ {
		if postorder[j] < postorder[right] {
			return false
		}
	}

	return verifyHelper(postorder, left, i-1) && verifyHelper(postorder, i, right-1)
}

// func main() {
// 	fmt.Println(verifyPostorder([]int{1, 6, 3, 2, 5}))
// 	fmt.Println(verifyPostorder([]int{1, 3, 2, 6, 5}))
// 	fmt.Println(verifyPostorder([]int{4, 8, 6, 12, 16, 14, 10}))
// 	fmt.Println(verifyPostorder([]int{4, 6, 7, 5}))
// 	fmt.Println(verifyPostorder([]int{4, 7, 6, 5}))
// 	fmt.Println(verifyPostorder([]int{1, 2, 3, 4, 5}))
// 	fmt.Println(verifyPostorder([]int{5, 4, 3, 2, 1}))
// 	fmt.Println(verifyPostorder([]int{5, 2, 3, 4, 1}))
// }
