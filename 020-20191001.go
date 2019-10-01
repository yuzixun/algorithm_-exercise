package main

func isValid(s string) bool {
	sMap := map[string]string{")": "(", "}": "{", "]": "["}
	sStack := []string{}
	for _, e := range s {
		es := string(e)
		if sMap[es] != "" {
			size := len(sStack)
			if size > 0 && sMap[es] == sStack[size-1] {
				sStack = sStack[:size-1]
			} else {
				return false
			}
		} else {
			sStack = append(sStack, es)
		}
	}

	return len(sStack) == 0
}

// func main() {
// 	// s := ")[]{}"
// 	// s := "([)]"
// 	// s := "{[]}"
// 	// s := "()[]{}"
// 	// s := "(]"

// 	// fmt.Println(isValid(s))
// }
