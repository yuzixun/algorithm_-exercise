package main

func partitionLabels(S string) []int {
	result := []int{}
	stringToPos := make(map[rune]int)

	start, last := 0, 0
	for i, l := range S {
		stringToPos[l] = i
	}
	for i, l := range S {
		if last < stringToPos[l] {
			last = stringToPos[l]
		}

		if last == i {
			result = append(result, last-start+1)
			start = last + 1
		}
	}

	// fmt.Println(result)
	return result
}

// func main() {
// 	partitionLabels("ababcbacadefegdehijhklij")
// }
