package main

type Trie struct {
	word     string
	children [26]*Trie
}

func findWords(board [][]byte, words []string) []string {
	root := &Trie{}
	for _, word := range words {
		node := root
		for _, item := range word {
			key := item - 'a'
			if node.children[key] == nil {
				node.children[key] = &Trie{}
			}
			node = node.children[key]
		}
		node.word = word
	}

	result := make([]string, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs4FindWords(board, i, j, root, &result)
		}
	}
	// fmt.Println(result)
	return result
}

func dfs4FindWords(board [][]byte, i, j int, node *Trie, result *[]string) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}
	c := board[i][j]
	if c == '#' || node.children[c-'a'] == nil {
		return
	}

	node = node.children[c-'a']
	if node.word != "" {
		*result = append(*result, node.word)
		node.word = ""
	}

	board[i][j] = '#'
	dfs4FindWords(board, i+1, j, node, result)
	dfs4FindWords(board, i-1, j, node, result)
	dfs4FindWords(board, i, j+1, node, result)
	dfs4FindWords(board, i, j-1, node, result)
	board[i][j] = c
}
func main() {
	findWords(
		[][]byte{
			[]byte{'o', 'a', 'a', 'n'},
			[]byte{'e', 't', 'a', 'e'},
			[]byte{'i', 'h', 'k', 'r'},
			[]byte{'i', 'f', 'l', 'v'}},
		[]string{"oath", "pea", "eat", "rain"},
	)
}
