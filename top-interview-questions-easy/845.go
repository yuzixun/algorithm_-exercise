package main

// func findClosestElements(arr []int, k int, x int) []int {
// 	left, right := 0, len(arr)-1
// 	removeCount := len(arr) - k
// 	for removeCount > 0 {
// 		if x-arr[left] > arr[right]-x {
// 			left++
// 		} else {
// 			right--
// 		}
// 		removeCount--
// 	}
// 	// fmt.Println(left, right, arr[left:right+1])
// 	return arr[left : right+1]
// }

func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k

	var mid int
	for left < right {
		mid = left + (right-left)/2
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// fmt.Println(arr[left : left+k])
	return arr[left : left+k]
}

func main() {
	findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3)
	findClosestElements([]int{1, 2, 3, 4, 5}, 4, -1)
	findClosestElements([]int{1, 2, 3, 4, 5, 6, 7}, 3, 5)
	findClosestElements([]int{0, 2, 2, 3, 4, 6, 7, 8, 9, 9}, 4, 5)
}
