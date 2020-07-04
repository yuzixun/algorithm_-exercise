package main

import (
	"sort"
)

type ByLen []string

func (a ByLen) Len() int {
	return len(a)
}

func (a ByLen) Less(i, j int) bool {
	return len(a[i]) > len(a[j])
}

func (a ByLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type Tree struct {
	// Val      byte
	Children [26]*Tree
}

func (this *Tree) insert(word string) int {
	new := false
	cur := this
	for i := len(word) - 1; i >= 0; i-- {
		c := word[i] - 'a'
		if cur.Children[c] == nil {
			new = true
			cur.Children[c] = &Tree{Children: [26]*Tree{}}
		}
		cur = cur.Children[c]
	}
	if new {
		return len(word) + 1
	}
	return 0
}
func minimumLengthEncoding(words []string) int {
	sort.Sort(ByLen(words))

	result := 0
	tree := &Tree{Children: [26]*Tree{}}
	for _, word := range words {
		result += tree.insert(word)
		// fmt.Println(tree)
	}

	// fmt.Println("result is ", result)
	return result
}

func main() {
	minimumLengthEncoding([]string{"time", "me", "bell"})
	minimumLengthEncoding([]string{"time", "me", "lme", "bell"})
}
