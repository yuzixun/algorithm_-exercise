package main

type Trie struct {
	word     string
	children [26]*Trie
}

func findWords(board [][]byte, words []string) []string {
	// 1. 建好树
	root := &Trie{}
	for _, word := range words {
		cur := root
		for _, item := range word {
			key := item - 'a'
			if cur.children[key] == nil {
				cur.children[key] = &Trie{}
			}
			cur = cur.children[key]
		}
		cur.word = word
	}

	rows, columns := len(board), len(board[0])
	result := []string{}
	// 2. 从根开始搜索单词
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			dfs(board, i, j, root, &result)
		}
	}

	return result
}

func dfs(board [][]byte, i, j int, node *Trie, result *[]string) {
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
	dfs(board, i+1, j, node, result)
	dfs(board, i, j+1, node, result)
	dfs(board, i-1, j, node, result)
	dfs(board, i, j-1, node, result)
	board[i][j] = c
}

func main() {

}
