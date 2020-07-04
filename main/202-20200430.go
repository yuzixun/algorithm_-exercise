package main

import "fmt"

// func isHappy(n int) bool {
// 	cache := map[int]bool{n: true}

// 	num := n
// 	for {
// 		num = getNumsSum(num)
// 		if num == 1 {
// 			return true
// 		}
// 		_, exist := cache[num]
// 		if exist {
// 			return false
// 		} else {

// 			cache[num] = true
// 		}
// 	}
// }
func isHappy(n int) bool {
	slow, fast := n, n
	for {
		slow = getNumsSum(slow)
		fast = getNumsSum(getNumsSum(fast))
		if fast == 1 {
			return true
		}
		if fast == slow {
			return false
		}
	}
}

func getNumsSum(n int) int {
	sum := 0

	var c int
	for n != 0 {
		c = (n % 10)
		sum += (c * c)
		n /= 10
	}

	return sum
}

func main() {
	fmt.Println(isHappy(19))
}
