package main

func divide(dividend int, divisor int) int {
	max, min := 1<<31-1, -1<<31

	if divisor == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}

	if dividend == min && divisor == -1 {
		return max
	}

	sign := 1
	if dividend > 0 && divisor < 0 {
		divisor = -1 * divisor
		sign = -1
	} else if dividend < 0 && divisor > 0 {
		dividend = -1 * dividend
		sign = -1
	} else if dividend < 0 && divisor < 0 {
		dividend = -1 * dividend
		divisor = -1 * divisor
	}

	count := 0
	for divisor <= dividend {
		count++
		divisor <<= 1
	}

	result := 0
	// fmt.Println("count is ", count)
	for count > 0 {
		count--
		divisor >>= 1
		temp := 1
		if dividend >= divisor {
			result += (temp << uint32(count))
			dividend -= divisor
		}
	}

	result *= sign

	if result > max {
		return max
	}
	if result < min {
		return min
	}
	// fmt.Println(result)
	return result
}

// func main() {
// 	divide(10, 3)
// 	divide(1, 1)
// 	divide(7, -3)
// 	divide(1, -1)
// 	divide(-1, -1)
// 	divide(-2147483648, -1)
// 	divide(2147483648, -1)
// }
