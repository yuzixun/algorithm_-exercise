package main

// func isPalindrome(x int) bool {
// 	if x < 0 || x%10 == 0 {
// 		return false
// 	}
// 	if x == 0 {
// 		return true
// 	}

// 	var arr []int
// 	i := 0
// 	for x != 0 {
// 		arr = append(arr, x%10)
// 		x /= 10
// 		i++
// 	}

// 	for left, right := 0, len(arr)-1; left <= right; {
// 		if arr[left] != arr[right] {
// 			return false
// 		}

// 		left++
// 		right--
// 	}

// 	return true
// }

// func isPalindrome(x int) bool {
// 	if x < 0 || (x != 0 && x%10 == 0) {
// 		return false
// 	}

// 	tmp, reversedNum := x, 0

// 	for tmp > 0 {
// 		reversedNum = (reversedNum*10 + tmp%10)
// 		tmp /= 10
// 	}
// 	fmt.Println(reversedNum)
// 	return reversedNum == x
// }

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	reversedNum := 0

	for x > reversedNum {
		c := x % 10
		x /= 10
		reversedNum = (reversedNum*10 + c)
	}
	return reversedNum == x || reversedNum/10 == x
}

// func main() {
// 	fmt.Println(isPalindrome(121))
// 	fmt.Println(isPalindrome(0))
// 	fmt.Println(isPalindrome(-121))
// 	fmt.Println(isPalindrome(10))
// }
