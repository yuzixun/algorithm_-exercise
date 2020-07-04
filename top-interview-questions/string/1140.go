package main

import "fmt"

type Trie struct {
	isEnd bool
	node  map[rune]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		isEnd: false,
		node:  map[rune]*Trie{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	curTrie := this

	for _, item := range word {
		trie := curTrie.node[item]
		if trie == nil {
			c := Constructor()
			trie = &c
			curTrie.node[item] = &c
		}
		curTrie = trie
	}
	curTrie.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	curTrie := this
	for _, item := range word {
		curTrie = curTrie.node[item]
		if curTrie == nil {
			return false
		}
	}
	return curTrie.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	curTrie := this
	for _, item := range prefix {
		curTrie = curTrie.node[item]
		if curTrie == nil {
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
/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
	trie := Constructor()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // 返回 true
	fmt.Println(trie.Search("app"))     // 返回 false
	fmt.Println(trie.StartsWith("app")) // 返回 true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // 返回 true
}
