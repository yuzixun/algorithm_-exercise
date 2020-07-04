package main

// I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
// X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
// C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

func romanToInt(s string) int {
	var prev byte
	result := 0

	for i := 0; i < len(s); i++ {

		switch s[i] {
		case 'I':
			result += 1

		case 'V':
			if prev == 'I' {
				result += 3
			} else {
				result += 5
			}

		case 'X':
			if prev == 'I' {
				result += 8
			} else {
				result += 10
			}

		case 'L':
			if prev == 'X' {
				result += 30
			} else {
				result += 50
			}

		case 'C':
			if prev == 'X' {
				result += 80
			} else {
				result += 100
			}

		case 'D':
			if prev == 'C' {
				result += 300
			} else {
				result += 500
			}

		case 'M':
			if prev == 'C' {
				result += 800
			} else {
				result += 1000
			}

		}
		prev = s[i]
	}
	// fmt.Println(result)
	return result
}

// func main() {
// 	romanToInt("III")
// 	romanToInt("IV")
// 	romanToInt("IX")
// 	romanToInt("LVIII")
// 	romanToInt("MCMXCIV")
// }
