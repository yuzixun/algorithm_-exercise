package main

func romanToInt(s string) int {
	result := 0
	size := len(s)
	for i := 0; i < size; i++ {
		switch s[i] {
		case 'M':
			result += 1000
		case 'D':
			result += 500
		case 'C':
			if i+1 < size && (s[i+1] == 'D' || s[i+1] == 'M') {
				if s[i+1] == 'D' {
					result += 400
					i++
				} else if s[i+1] == 'M' {
					result += 900
					i++
				}
			} else {
				result += 100
			}
		case 'L':
			result += 50
		case 'X':
			if i+1 < size && (s[i+1] == 'L' || s[i+1] == 'C') {
				if s[i+1] == 'L' {
					result += 40
					i++
				} else if s[i+1] == 'C' {
					result += 90
					i++
				}
			} else {
				result += 10
			}
		case 'V':
			result += 5
		case 'I':
			if i+1 < size && (s[i+1] == 'V' || s[i+1] == 'X') {
				if s[i+1] == 'V' {
					result += 4
					i++
				} else if s[i+1] == 'X' {
					result += 9
					i++
				}
			} else {
				result++
			}
		}
	}

	// fmt.Println(result)
	return result
}

func main() {
	romanToInt("III")
	romanToInt("IV")
	romanToInt("IX")
	romanToInt("LVIII")
	romanToInt("MCMXCIV")
}
