package main

import "fmt"

type Trie struct {
	isEnd bool
	node  map[rune]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	node := make(map[rune]*Trie)
	trie := Trie{node: node, isEnd: false}
	return trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	prev := this
	for _, s := range word {
		trie := prev.node[s]
		if trie == nil {
			c := Constructor()
			trie = &c
			prev.node[s] = trie
		}
		prev = trie
	}
	prev.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	c := this
	for _, s := range word {
		trie := c.node[s]
		if trie != nil {
			c = trie
		} else {
			return false
		}
	}
	return c.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	c := this
	for _, s := range prefix {
		trie := c.node[s]
		if trie != nil {
			c = trie
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	obj := Constructor()

	obj.Insert("app")
	obj.Insert("apple")
	obj.Insert("beer")
	obj.Insert("add")
	obj.Insert("jam")
	obj.Insert("rental")
	// fmt.Println(obj.Search("apps"))
	// fmt.Println(obj.Search("app"))
	// fmt.Println(obj.Search("ad"))
	// fmt.Println(obj.Search("applepie"))
	// fmt.Println(obj.Search("rest"))
	// fmt.Println(obj.Search("jan"))
	// fmt.Println(obj.Search("rent"))
	// fmt.Println(obj.Search("beer"))
	// fmt.Println(obj.Search("jam"))
	// fmt.Println(obj.StartsWith("apps"))
	// fmt.Println(obj.StartsWith("app"))
	fmt.Println(obj.StartsWith("ad")) // e
	// fmt.Println(obj.StartsWith("applepie"))
	// fmt.Println(obj.StartsWith("rest"))
	// fmt.Println(obj.StartsWith("jan"))
	fmt.Println(obj.StartsWith("rent")) // e
	// fmt.Println(obj.StartsWith("beer"))
	// fmt.Println(obj.StartsWith("jam"))

}
