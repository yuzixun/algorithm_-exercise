package main

import "fmt"

func equationsPossible(equations []string) bool {
	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
		parent[i] = i
	}

	for _, str := range equations {
		if str[1] == '=' {
			union(parent, int(str[0]-'a'), int(str[3]-'a'))
		}
	}

	// fmt.Println(parent)
	for _, str := range equations {
		if str[1] == '!' {
			if find(parent, int(str[0]-'a')) == find(parent, int(str[3]-'a')) {
				return false
			}
		}
	}

	return true
}

func union(parent []int, a, b int) {
	parent[find(parent, a)] = find(parent, b)
}

func find(parent []int, index int) int {
	for parent[index] != index {
		parent[index] = parent[parent[index]]
		index = parent[index]
	}
	// fmt.Println(index, parent[index])
	return index
}

func main() {
	fmt.Println(equationsPossible([]string{"a==b", "b!=a"}))
	fmt.Println(equationsPossible([]string{"b==a", "a==b"}))
	fmt.Println(equationsPossible([]string{"a==b", "b==c", "a==c"}))
	fmt.Println(equationsPossible([]string{"a==b", "b!=c", "c==a"}))
	fmt.Println(equationsPossible([]string{"c==c", "b==d", "x!=z"}))
}
